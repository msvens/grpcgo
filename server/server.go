package main

import (
	"context"
	"fmt"
	"github.com/msvens/grpcgo/db"
	pb "github.com/msvens/grpcgo/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

const (
	port = ":50051"
	)

type server struct {
	pb.UnimplementedUserServiceServer
	dao db.UserDao
}

func (s *server) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	id, err := s.dao.Create(req)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.AddUserResponse{Id: id}, nil
}

func (s *server) DeleteAll(ctx context.Context, req *pb.DeleteAllRequest) (*pb.DeleteAllResponse, error) {
	count, err := s.dao.DeleteAll()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.DeleteAllResponse{Count:int32(count)}, nil
}

func (s *server) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	count, err := s.dao.Delete(req)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return &pb.DeleteUserResponse{Count: int32(count)}, nil
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	u, err := s.dao.Get(req.Id)
	switch err {
	case nil:
		return u, nil
	case db.NoSuchUser:
		return nil, status.Error(codes.NotFound, err.Error())
	default:
		return nil, status.Error(codes.Internal, err.Error())
	}
}

func (s *server) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserResponse, error) {
	users, err := s.dao.List()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	return users, nil
}

func startDb() (*db.PgUserDao, error) {
	fmt.Println("starting db and creating tables")
	dao, err := db.NewPgDao()
	if err != nil {
		return nil, err
	}
	err = dao.CreateTables()
	return dao, err
}

func main() {

	dao,err := startDb()

	if err != nil {
		panic(err)
	}
	defer dao.Close()

	//start server
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s :=  grpc.NewServer()
	pb.RegisterUserServiceServer(s, &server{dao:dao})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	log.Printf("signal signal")


}