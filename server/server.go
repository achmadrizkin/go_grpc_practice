package server

import (
	"context"
	"errors"
	"go_grpc_practice/model"
	"go_grpc_practice/service"
	pb "go_grpc_practice/v1_book_grpc"
)

type Server struct {
}

func mapToBookpb(book model.Book) *pb.Book {
	return &pb.Book{
		Id:     book.Id,
		Author: book.Author,
		Title:  book.Title,
		IsRead: book.IsRead,
	}
}

// mendapatkan data berdasarkan id
func (*Server) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.GetBookResponse, error) {
	var bookId string = req.Book.GetId()
	_, result := service.GetBook(bookId)

	// jika tidak ditemukan, show error
	if result.Id != bookId {
		return &pb.GetBookResponse{Status: false, Data: nil}, errors.New("Data not found!")
	}

	// mengolah hasil dari GetBook
	// untuk data di objek response
	var bookData *pb.Book = &pb.Book{
		Id:     result.Id,
		Author: result.Author,
		Title:  result.Title,
		IsRead: result.IsRead,
	}

	// memberikan response berupa data buku
	return &pb.GetBookResponse{
		Status: true,
		Data:   bookData,
	}, nil
}
