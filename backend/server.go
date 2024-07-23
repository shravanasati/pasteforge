package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"github.com/shravanasati/pasteforge/backend/db"
	"github.com/shravanasati/pasteforge/backend/services/misc"
	"github.com/shravanasati/pasteforge/backend/services/pastes"
)

// env vars
var (
	PORT              string
	GIN_MODE          string
	DIST_DIR          string
	SECRET_KEY        string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_DB       string
)

func validateNotEmpty(key, value string) {
	if value == "" {
		panic(key + " env var not configured")
	}
}

func validatePort(val string) {
	matches, err := regexp.MatchString(`^\d{4,5}$`, val)
	if err != nil {
		panic("error validating port" + err.Error())
	}
	if !matches {
		panic(fmt.Sprintf("env var PORT=%s is incorrect", val))
	}
}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("unable to load env variables")
	}

	PORT = os.Getenv("PORT")
	validateNotEmpty("PORT", PORT)
	validatePort(PORT)

	GIN_MODE = os.Getenv("GIN_MODE")
	validateNotEmpty("GIN_MODE", GIN_MODE)

	DIST_DIR = os.Getenv("DIST_DIR")
	validateNotEmpty("DIST_DIR", DIST_DIR)

	SECRET_KEY = os.Getenv("SECRET_KEY")
	validateNotEmpty("SECRET_KEY", SECRET_KEY)

	POSTGRES_USER = os.Getenv("POSTGRES_USER")
	validateNotEmpty("POSTGRES_USER", POSTGRES_USER)
	POSTGRES_PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	validateNotEmpty("POSTGRES_PASSWORD", POSTGRES_PASSWORD)
	POSTGRES_DB = os.Getenv("POSTGRES_DB")
	validateNotEmpty("POSTGRES_DB", POSTGRES_DB)
}

func initDB() *pgxpool.Pool {
	conn, err := db.NewConnPool(POSTGRES_USER, POSTGRES_PASSWORD, POSTGRES_DB)
	if err != nil {
		panic(err)
	}
	return conn
}

func main() {
	db := initDB()
	defer db.Close()

	gin.SetMode(GIN_MODE)
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))

	apiRouter := router.Group("/api")
	v1Router := apiRouter.Group("/v1")

	misc.RegisterRoutes(v1Router)
	pastesHandler := pastes.NewHandler(db)
	pastesHandler.RegisterRoutes(v1Router)

	router.NoRoute(gin.WrapH(http.FileServer(http.Dir(DIST_DIR))))

	server := &http.Server{
		Addr:           ":" + PORT,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20, // 1MB
	}

	log.Fatal(server.ListenAndServe())
}
