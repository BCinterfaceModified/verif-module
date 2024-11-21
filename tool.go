package verif_module

import (
	"fmt"
	"log"
	"net"

	pb "github.com/BCinterfaceModified/verif-module/bcinterface"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedNetworkInterfaceServer
}

var portNumber = "30001"

func startGrpcServer() {
	lis, err := net.Listen("tcp", portNumber)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterNetworkInterfaceServer(s, &Server{})
	fmt.Println("*************************************************************")
	fmt.Println("*                                                           *")
	fmt.Println("*                Running Commu Module GRPC                  *")
	fmt.Println("*                                                           *")
	fmt.Println("*************************************************************")

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
