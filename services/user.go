package services

import (
	
	"fmt"
	"context"
	"time"
	"log"
	"io"
	"github.com/ramoncgusmao/test-grpc-go/pb"
)
func NewUserService() *UserService {
	return &UserService{}
}
type UserService struct {
	pb.UnimplementedUserServiceServer
}
func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	
	fmt.Println(req.Name)
	return &pb.User{
			Id: "123",
			Name: req.GetName(),
			Email: req.GetEmail(),
		}, nil
}

func (*UserService) AddUserVerbose(req *pb.User, stream pb.UserService_AddUserVerboseServer) error {

	stream.Send(&pb.UserResultStream{
		Status: "Init",
		User: &pb.User{},
	})

	time.Sleep(time.Second * 3)

	stream.Send(&pb.UserResultStream{
		Status: "Insert",
		User: &pb.User{},
	})
	time.Sleep(time.Second * 3)
	
	stream.Send(&pb.UserResultStream{
		Status: "User has been inserted",
		User: &pb.User{
			Id: "123",
			Name: req.GetName(),
			Email: req.GetEmail(),
		},
	})

	time.Sleep(time.Second * 3)
	
	stream.Send(&pb.UserResultStream{
		Status: "Completed",
		User: &pb.User{
			Id: "123",
			Name: req.GetName(),
			Email: req.GetEmail(),
		},
	})

	return nil

}

func (*UserService) AddUsers(stream pb.UserService_AddUsersServer) error {
	
	users := []*pb.User{}

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(
				&pb.Users{
					User: users,
				})
		}

		if err != nil {
			log.Fatalf("Error receiving stream: %v", err)
		}

		users = append(users, &pb.User{
			Id: req.GetId(),
			Name: req.GetName(),
			Email: req.GetEmail(),
		})

		fmt.Println("Adding", req.GetName())
	}
}


func (*UserService) AddUsersStreamBoth(stream pb.UserService_AddUsersStreamBothServer) error{

	for{
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error receiving stream: %v", err)
		}

		err = stream.Send(&pb.UserResultStream{
			Status: "Added",
			User: req,
		})
		if err != nil {
			log.Fatalf("Error sended stream: %v", err)
		}
	}
}


/*
type UserServiceServer interface {
	AddUser(context.Context, *User) (*User, error)
	AddUserVerbose(*User, UserService_AddUserVerboseServer) error
	mustEmbedUnimplementedUserServiceServer()
}
*/

