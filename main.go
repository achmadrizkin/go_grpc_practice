package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"go_grpc_practice/server"
	pb "go_grpc_practice/v1_book_grpc"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// implementasi rpc dengan nama GetBook
//
// -- GRPC 1 Practice
// func (*server) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.GetBookResponse, error) {
// 	// menerima request
// 	bookReq := req.Book

// 	// mengolah hasil dari request
// 	book := pb.Book{
// 		Id:     bookReq.Id,
// 		Title:  bookReq.Title,
// 		Author: bookReq.Author,
// 		IsRead: false,
// 	}

// 	// mengirim respons
// 	return &pb.GetBookResponse{
// 		Status: true,
// 		Data:   &book,
// 	}, nil

// }

func main() {
	// jika kode mengalami crash, nomor line akan ditampilkan
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Book service started")

	// membuat gRPC server
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	s := grpc.NewServer()
	pb.RegisterBookServiceServer(s, &server.Server{})
	reflection.Register(s)

	go func() {
		fmt.Println("Starting server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Menunggu hingga dihentikan dengan Ctrl + C
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Lakukan block hingga sinyal sudah didapatkan
	<-ch
	fmt.Println("Stopping the server..")
	s.Stop()
	fmt.Println("Stopping listener...")
	lis.Close()
	fmt.Println("End of Program")
}
