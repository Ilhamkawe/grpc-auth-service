package grpc

import (
	"fmt"
	"github.com/Ilhamkawe/grpc-auth-service/internal/port"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GrpcAdapter struct {
	jwtService  port.JwtServicePort
	authService port.AuthServicePort
	grpcPort    int
	server      *grpc.Server
}

func NewGrpcAdapter(jwtService port.JwtServicePort, authService port.AuthServicePort, grpcPort int) *GrpcAdapter {
	return &GrpcAdapter{
		jwtService:  jwtService,
		authService: authService,
		grpcPort:    grpcPort,
	}
}

func (a *GrpcAdapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.grpcPort))

	if err != nil {
		log.Fatalf("Failed to listen on port :#{a.grpcPort} : #{err}\n")
	}

	log.Printf("Server listening on port : #{a.grpcPort} ")

	grpcServer := grpc.NewServer()
	a.server = grpcServer

	//register service grpc disini
	//auth.RegisterAuthServiceServer(grpcServer, a)

	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve on port #{a.grpcPort} : #{err}\n")
	}
}
