package main

import (
	"github.com/drTragger/semiTrashAPI/utils"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

const (
	apiPrefix string = "/api/v1"
)

var (
	port                string
	bookResourcePrefix  = apiPrefix + "/book" // /api/v1/book/
	booksResourcePrefix = apiPrefix + "/books"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Could not find .env file")
	}
	port = os.Getenv("app_port")
}

func main() {
	log.Println("Starting REST API server on port:", port)
	router := mux.NewRouter()

	utils.BuildBookResource(router, bookResourcePrefix)
	utils.BuildManyBooksResourcePrefix(router, booksResourcePrefix)

	log.Println("Router has been initialized successfully. Ready to go!")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
