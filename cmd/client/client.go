package main

import ("google.golang.org/grpc"
		"github.com/ramoncgusmao/test-grpc-go/pb"
		"log"
		"context"
		"fmt"
		"io"
		"time"

		)


func main(){
	connection, err := grpc.Dial("localhost:50051",grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect gRPC Server: %v", err)
	}
	defer connection.Close()
	
	client := pb.NewUserServiceClient(connection)
	//AddUser(client)
	//AddUserVerbose(client)
	//AddUsers(client)
	AddUsersStreamBoth(client)
}

func AddUser(client pb.UserServiceClient ){

	req := &pb.User {
		Id: "0",
		Name: "Ramon",
		Email: "j0j@com",
	}

	response, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}
	fmt.Println(response)

}


func AddUserVerbose(client pb.UserServiceClient ){

	req := &pb.User {
		Id: "0",
		Name: "Ramon",
		Email: "j0j@com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for{
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receive the msg: %v", err)
		}
		fmt.Println("Status", stream.Status)
	}
}


func AddUsers(client pb.UserServiceClient ){

	reqs := []* pb.User {
		&pb.User {
			Id: "0",
			Name: "Ramon",
			Email: "j0j@com",
		},
		&pb.User {
			Id: "1",
			Name: "Ramon1",
			Email: "j0j@com",
		},
		&pb.User {
			Id: "2",
			Name: "Ramon2",
			Email: "j0j@com",
		},
		&pb.User {
			Id: "3",
			Name: "Ramon3",
			Email: "j0j@com",
		},
		&pb.User {
			Id: "4",
			Name: "Ramon4",
			Email: "j0j@com",
		},
		&pb.User {
			Id: "5",
			Name: "Ramon5",
			Email: "j0j@com",
		},
		&pb.User {
			Id: "6",
			Name: "Ramon6",
			Email: "j0j@com",
		},	
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("error creating request: %v", err)
	}

	for _, req := range reqs{
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error creating request: %v", err)
	}

	fmt.Println(res)

}

func AddUsersStreamBoth(client pb.UserServiceClient ){

	stream, err := client.AddUsersStreamBoth(context.Background())
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}

	reqs := []* pb.User {
		&pb.User {
			Id: "0",
			Name: "Ramon",
			Email: "j0j@com",
		},
		&pb.User {
			Id: "1",
			Name: "Ramon1",
			Email: "j0j@com",
		},
		&pb.User {
			Id: "2",
			Name: "Ramon2",
			Email: "j0j@com",
		},
		&pb.User {
			Id: "3",
			Name: "Ramon3",
			Email: "j0j@com",
		},
		&pb.User {
			Id: "4",
			Name: "Ramon4",
			Email: "j0j@com",
		},
		&pb.User {
			Id: "5",
			Name: "Ramon5",
			Email: "j0j@com",
		},
		&pb.User {
			Id: "6",
			Name: "Ramon6",
			Email: "j0j@com",
		},	
	}

	wait := make(chan int)
	go func() {
		for _, req := range reqs {
			fmt.Println("Sending user: ", req.Name)
			stream.Send(req)
			time.Sleep(time.Second * 2)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			
			if err != nil {
				log.Fatalf("Error receiving data: %v",err)
			}
			fmt.Printf("Recebendo user %v como o status: %v\n", res.GetUser().GetName(), res.GetStatus())

		}
		close(wait)
	}()

	<- wait
}