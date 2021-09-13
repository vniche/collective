package main

import (
	"flag"

	_ "github.com/joho/godotenv/autoload"
	"github.com/vniche/collective/server/processor"
	"github.com/vniche/collective/server/services"
)

func main() {
	workload := flag.String("workload", "grpc", "--workload=processor or --workload=grpc-web")
	flag.Parse()

	switch *workload {
	case "processor":
		processor.Start()
	case "grpc-web":
		services.StartGrpcWeb()
	default:
		services.StartGrpc()
	}
}
