package main

import (
	"database/sql"
	myDb "github.com/Ilhamkawe/grpc-auth-service/internal/adapter/database"
	mygrpc "github.com/Ilhamkawe/grpc-auth-service/internal/adapter/grpc"
	app "github.com/Ilhamkawe/grpc-auth-service/internal/application/application"
	"log"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(logWriter{})

	sqlDB, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5432/grpc?sslmode=disable")

	if err != nil {
		log.Fatalln("Can't connect database :", err)
	}

	databaseAdapter, err := myDb.NewDatabaseAdapter(sqlDB)

	if err != nil {
		log.Fatalln("Can't create database adapter :", err)
	}

	as := app.NewAuthService(databaseAdapter)
	js := app.NewJwtService()

	grpcAdapter := mygrpc.NewGrpcAdapter(js, as, 9090)

}
