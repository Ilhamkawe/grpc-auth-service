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

	sqlDB, err := sql.Open("pgx", "postgres://postgres:kawe123@localhost:5433/crowdfunding?sslmode=disable")

	if err != nil {
		log.Fatalln("Can't connect database :", err)
	}

	databaseAdapter, err := myDb.NewDatabaseAdapter(sqlDB)

	if err != nil {
		log.Fatalln("Can't create database adapter :", err)
	}

	rdbAdapter, err := myDb.NewJwtRedisAdapter("localhost:6379", "", "", 0)

	if err != nil {
		log.Fatalln("Can't create Redis database adapter :", err)
	}

	as := app.NewAuthService(databaseAdapter, rdbAdapter)
	js := app.NewJwtService()
	fs := app.NewFileService()

	grpcAdapter := mygrpc.NewGrpcAdapter(rdbAdapter, fs, js, as, 9090)

	grpcAdapter.Run()
}
