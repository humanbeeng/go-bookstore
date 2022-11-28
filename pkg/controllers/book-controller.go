package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/humanbeeng/go-bookstore/pkg/models"
	"github.com/humanbeeng/go-bookstore/pkg/utils"
)

var NewBook models.Book

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing bookId")
	}

	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r, book)
	b := book.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing bookId")
	}

	b := models.DeleteBook(ID)
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var book = &models.Book{}

	utils.ParseBody(r, book)

	vars := mux.Vars(r)

	bookId := vars["bookId"]

	ID, err := strconv.ParseInt(bookId, 0, 0)

	if err != nil {
		fmt.Println("error while parsing")
	}

	bookDetails, db := models.GetBookById(ID)
	models.DeleteBook(int64(bookDetails.ID))

	book.Name = bookDetails.Name
	book.Author = bookDetails.Author
	book.Publication = bookDetails.Publication

	db.Save(&book)
	res, _ := json.Marshal(book)

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
