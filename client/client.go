package main

import (
	"log"
	"context"
	//"os"
	"time"
	"google.golang.org/grpc"
	pb "github.com/msvens/grpcgo/service"
)
const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 2)
	defer cancel()

	cnt, err := c.DeleteAll(ctx, &pb.DeleteAllRequest{})
	if err != nil {
		log.Fatalf("could not delete all users: %v", err)
	}
	log.Printf("deleted users: %v", cnt)

	r, err := c.AddUser(ctx, &pb.AddUserRequest{Name: "Martin Svensson", Email: "msvens3@gmail.com"})
	if err != nil {
		log.Fatalf("could not add user: %v", err)
	}
	log.Printf("Added user: %v", r)

	l, err := c.ListUser(ctx, &pb.ListUserRequest{})
	if err != nil {
		log.Fatalf("could not list users: %v", err)
	}
	for idx,u := range l.Users {
		log.Printf("%v: %v\n", idx, u)
	}

}