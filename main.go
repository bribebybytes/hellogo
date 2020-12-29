package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	loadEnvFile()
}

func loadEnvFile() {
	log.Print("Loading env variables")
	if err := godotenv.Load(os.Getenv("ENVVARS_PATH") + "/.env"); err != nil {
		log.Print("No .env file found")
	} else {
		log.Print(".Env file found. parsing env variables...")
		AppUname, _ = os.LookupEnv("APP_UNAME")
		AppPwd, _ = os.LookupEnv("APP_PWD")
		AppExtEndpoint, _ = os.LookupEnv("APP_EXT_ENDPOINT")
	}
}

func main() {

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
