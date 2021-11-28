package handlers

import (
	"encoding/json"
	"github.com/drTragger/semiTrashAPI/models"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

func GetBookById(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("Error while parsing happened:", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "Do not use parameter ID which cannot be casted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	book, ok := models.FindBookById(id)
	log.Println("Get book with id:", id)
	if !ok {
		writer.WriteHeader(404)
		msg := models.Message{Message: "Book with this id does not exist"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	writer.WriteHeader(200)
	json.NewEncoder(writer).Encode(book)
}

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Create new book")
	var book models.Book

	err := json.NewDecoder(request.Body).Decode(&book)
	if err != nil {
		msg := models.Message{Message: "Provided JSON file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	book.ID = len(models.DB) + 1
	models.DB = append(models.DB, book)

	writer.WriteHeader(201)
	json.NewEncoder(writer).Encode(book)
}

func UpdateBook(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Update the book")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("Error while parsing happened:", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "Do not use parameter ID which cannot be casted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	oldBook, ok := models.FindBookById(id)
	var newBook models.Book
	if !ok {
		log.Println("Book not found. id:", id)
		writer.WriteHeader(404)
		msg := models.Message{Message: "Book with this id does not exist"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	err = json.NewDecoder(request.Body).Decode(&newBook)
	if err != nil {
		msg := models.Message{Message: "Provided JSON file is invalid"}
		writer.WriteHeader(400)
		json.NewEncoder(writer).Encode(msg)
		return
	}

	for i, b := range models.DB {
		if b.ID == oldBook.ID {
			models.DB[i] = newBook
			models.DB[i].ID = oldBook.ID
		}
	}

	writer.WriteHeader(200)
	msg := models.Message{Message: "Updated the book"}
	json.NewEncoder(writer).Encode(msg)
}

func DeleteBook(writer http.ResponseWriter, request *http.Request) {
	initHeaders(writer)
	log.Println("Delete the book")
	id, err := strconv.Atoi(mux.Vars(request)["id"])
	if err != nil {
		log.Println("Error while parsing happened:", err)
		writer.WriteHeader(400)
		msg := models.Message{Message: "Do not use parameter ID which cannot be casted to int type"}
		json.NewEncoder(writer).Encode(msg)
		return
	}
	book, ok := models.FindBookById(id)
	if !ok {
		log.Println("Book not found. id:", id)
		writer.WriteHeader(404)
		msg := models.Message{Message: "Book with this id does not exist"}
		json.NewEncoder(writer).Encode(msg)
		return
	}

	for i, b := range models.DB {
		if b.ID == book.ID {
			models.DB = append(models.DB[:i], models.DB[i+1:]...)
			break
		}
	}

	msg := models.Message{Message: "The book has been successfully deleted"}
	json.NewEncoder(writer).Encode(msg)
}
