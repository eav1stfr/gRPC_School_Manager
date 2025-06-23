package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"grpcapi/internals/api/handlers"
	mongo "grpcapi/internals/repositories/mongodb"
	"log"
	"net"
	"os"
)
import pb "grpcapi/proto/gen"

func main() {
	_, err := mongo.CreateMongoClient()
	if err != nil {
		fmt.Println(err)
	}
	err = godotenv.Load()
	if err != nil {
		log.Fatal("err loading .env file:", err)
	}

	s := grpc.NewServer()
	pb.RegisterExecsServiceServer(s, &handlers.Server{})
	pb.RegisterStudentsServiceServer(s, &handlers.Server{})
	pb.RegisterTeachersServiceServer(s, &handlers.Server{})

	reflection.Register(s)

	port := os.Getenv("SERVER_PORT")
	fmt.Println("grpc server is running on port 50051")
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("Error listening on port  50051")
	}
	err = s.Serve(lis)

	if err != nil {
		log.Fatal("Filed to serve:", err)
	}
}
