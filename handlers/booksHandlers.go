package handlers

import (
	"encoding/json"
	"github.com/drTragger/semiTrashAPI/models"
	"log"
	"net/http"
)

func initHeaders(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
}

func GetAllBooks(writer http.ResponseWriter, response *http.Request) {
	initHeaders(writer)
	log.Println("Get all books from DataBase")
	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(models.DB)
}
