package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

// env vars
var (
	ADDR              string
	PORT              string
	GIN_MODE          string
	DIST_DIR          string
	SECRET_KEY        string
	POSTGRES_USER     string
	POSTGRES_PASSWORD string
	POSTGRES_HOSTNAME string
	POSTGRES_PORT     string
	POSTGRES_DB       string
)

func validateNotEmpty(key, value string) {
	if value == "" {
		panic(key + " env var not configured")
	}
}

func validatePort(key, val string) {
	matches, err := regexp.MatchString(`^\d{4,5}$`, val)
	if err != nil {
		panic("error validating port" + err.Error())
	}
	if !matches {
		panic(fmt.Sprintf("env var %s=%s is incorrect", key, val))
	}
}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("unable to load env variables")
	}

	ADDR = os.Getenv("ADDR")
	validateNotEmpty("ADDR", ADDR)

	PORT = os.Getenv("PORT")
	validateNotEmpty("PORT", PORT)
	validatePort("PORT", PORT)

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
	POSTGRES_HOSTNAME = os.Getenv("POSTGRES_HOSTNAME")
	validateNotEmpty("POSTGRES_HOSTNAME", POSTGRES_HOSTNAME)
	POSTGRES_PORT = os.Getenv("POSTGRES_PORT")
	validateNotEmpty("POSTGRES_PORT", POSTGRES_PORT)
	validatePort("POSTGRES_PORT", POSTGRES_PORT)
}
