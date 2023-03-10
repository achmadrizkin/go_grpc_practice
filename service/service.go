package service

import (
	"go_grpc_practice/model"
	"go_grpc_practice/repository"
)

func AddBook(bookData model.Book) model.Book {
    return repository.AddBook(bookData)
}

func GetBook(bookId string) (int, model.Book) {
    return repository.GetBook(bookId)
}

func UpdateBook(bookData model.Book, id string) model.Book {
    return repository.UpdateBook(bookData, id)
}

func DeleteBook(id string) bool {
    return repository.DeleteBook(id)
}