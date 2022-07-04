package main
import (
	"log"
	"net"
	"github.com/ramoncgusmao/test-grpc-go/services"
	"google.golang.org/grpc"
	"github.com/ramoncgusmao/test-grpc-go/pb"
	"google.golang.org/grpc/reflection"
)

func main(){

	lis, err := net.Listen("tcp","localhost:50051");
	if err != nil {
		log.Fatalf("Could not connect %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())
	
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Could not server: %v", err)
	}
}