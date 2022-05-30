package main

import (
	"backend/internal/adapters/app/api"
	"backend/internal/adapters/core/arithmetic"
	rpc "backend/internal/adapters/framework/left/grpc"
	"backend/internal/adapters/framework/right/db"
	"backend/internal/ports"
	"log"
	"os"
)

func main() {
	var err error

	// ports
	var dbaseAdapter ports.DbPort
	var core ports.ArtithmeticPort
	var appAdapter ports.APIPort
	var grpcAdapter ports.GRPCPort

	dbaseDriver := os.Getenv("DB_DRIVER")
	dsourceName := os.Getenv("DS_NAME")
	dbaseAdapter, err = db.NewAdapter(dbaseDriver, dsourceName)
	if err != nil {
		log.Fatalf("failed to initiate dbase connection: %v", err)
	}
	defer dbaseAdapter.CloseDbConnection()

	core = arithmetic.NewAdapter()

	appAdapter = api.NewAdapter(dbaseAdapter, core)

	grpcAdapter = rpc.NewAdapter(appAdapter)
	grpcAdapter.Run()
}
