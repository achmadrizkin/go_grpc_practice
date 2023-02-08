package repository

import (
	"go_grpc_practice/model"
)

// membuat penyimpanan local dengan menggunakan slice
var storage []model.Book = []model.Book{}

func AddBook(bookData model.Book) model.Book {
	storage = append(storage, bookData)
	return bookData
}

func GetBook(bookId string) (int, model.Book) {
	for index, v := range storage {
		if v.Id == bookId {
			return index, v
		}
	}

	return 0, model.Book{}
}

func UpdateBook(bookData model.Book, id string) model.Book {
	index, book := GetBook(id)

	book.Title = bookData.Title
	book.Author = bookData.Author
	book.IsRead = bookData.IsRead

	storage[index] = book

	return book
}

func DeleteBook(id string) bool {
	var afterDeleted []model.Book = []model.Book{}
	for _, v := range storage {
		if id != v.Id {
			afterDeleted = append(afterDeleted, v)
		}
	}

	storage = afterDeleted
	return true
}
