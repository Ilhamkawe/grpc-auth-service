package grpc

import (
	"fmt"
	"github.com/Ilhamkawe/crowdfunding-proto/protogen/go/auth"
	"github.com/Ilhamkawe/grpc-auth-service/internal/interceptor"
	"github.com/Ilhamkawe/grpc-auth-service/internal/port"
	"github.com/dimk00z/grpc-filetransfer/pkg/logger"
	"google.golang.org/grpc"
	"log"
	"net"
)

type GrpcAdapter struct {
	jwtService  port.JwtServicePort
	authService port.AuthServicePort
	fileService port.FileServicePort
	rdb         port.JwtRedisPort
	log         *logger.Logger
	grpcPort    int
	server      *grpc.Server
}

func NewGrpcAdapter(rdb port.JwtRedisPort, fileService port.FileServicePort, jwtService port.JwtServicePort, authService port.AuthServicePort, grpcPort int) *GrpcAdapter {
	return &GrpcAdapter{
		rdb:         rdb,
		jwtService:  jwtService,
		authService: authService,
		fileService: fileService,
		grpcPort:    grpcPort,
	}
}

func (a *GrpcAdapter) Run() {
	var err error
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", a.grpcPort))

	if err != nil {
		log.Fatalf("Failed to listen on port :#{a.grpcPort} : #{err}\n")
	}

	log.Printf("Server listening on port : %d", a.grpcPort)

	// register method yang ingin dipasang interceptor
	protectedMethods := map[string]bool{
		"/auth.AuthService/FetchUser":      true,
		"/auth.AuthService/ChangePassword": true,
		"/auth.AuthService/UpdateUserInfo": true,
		"/auth.AuthService/Logout":         true,
	}

	//register interceptor
	authInterceptor := interceptor.AuthUnaryInterceptor(a.rdb, a.authService, a.jwtService, protectedMethods)

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(authInterceptor))
	a.server = grpcServer

	//register service grpc disini
	auth.RegisterAuthServiceServer(grpcServer, a)

	if err = grpcServer.Serve(listen); err != nil {
		log.Fatalf("Failed to serve on port #{a.grpcPort} : #{err}\n")
	}
}
