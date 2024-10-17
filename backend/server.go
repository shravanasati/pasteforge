package main

import (
	"context"
	"errors"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/shravanasati/pasteforge/backend/db"
	"github.com/shravanasati/pasteforge/backend/services/misc"
	"github.com/shravanasati/pasteforge/backend/services/pastes"
)

func initDB() *pgxpool.Pool {
	conn, err := db.NewConnPool(POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_HOSTNAME, POSTGRES_PORT, POSTGRES_DB)
	if err != nil {
		panic(err)
	}
	return conn
}

func main() {
	db := initDB()
	defer db.Close()

	logger := slog.Default()

	gin.SetMode(GIN_MODE)
	router := gin.Default()
	defaultRateLimiter := NewRateLimiter(20, time.Second)
	apiRateLimiter := NewRateLimiter(5, time.Second)
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	apiRouter := router.Group("/api")
	apiRouter.Use(apiRateLimiter)
	v1Router := apiRouter.Group("/v1")

	misc.RegisterRoutes(v1Router)

	minuteTicker := time.NewTicker(time.Minute)
	defer minuteTicker.Stop()
	done := make(chan struct{})
	pastesHandler := pastes.NewHandler(logger, db)
	go pastesHandler.DeleteExpiredPastes(minuteTicker, done)
	pastesHandler.RegisterRoutes(v1Router)

	router.Use(defaultRateLimiter, static.Serve("/", static.LocalFile(DIST_DIR, true)))
	router.NoRoute(defaultRateLimiter, func(ctx *gin.Context) {
		if !strings.HasPrefix(ctx.Request.RequestURI, "/api") {
			ctx.File(DIST_DIR + "/index.html")
			return
		}

		ctx.AbortWithStatus(http.StatusNotFound)
	})

	server := &http.Server{
		Addr:           ADDR + ":" + PORT,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		logger.Info("Listening at", "port", PORT)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal(err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Info("Shutting down server...")

	// send a signal to DeleteExpiredPastes goroutine to stop
	done <- struct{}{}

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	logger.Info("Server exiting...")
}
