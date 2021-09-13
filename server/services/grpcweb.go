package services

import (
	"log"
	"net/http"
	"os"

	"github.com/improbable-eng/grpc-web/go/grpcweb"
)

func StartGrpcWeb() {
	service := New()

	grpcWebPort, ok := os.LookupEnv("GRPC_WEB_PORT")
	if !ok {
		grpcWebPort = "0.0.0.0:8080"
	}

	wrappedGrpc := grpcweb.WrapServer(service.server)

	corsOrigin := "*"

	devMode, ok := os.LookupEnv("DEV_MODE")
	if !ok || devMode == "false" {
		corsOriginEnv, ok := os.LookupEnv("CORS_ORIGIN")
		if ok {
			corsOrigin = corsOriginEnv
		}
	}

	log.Printf("server listening at %v", grpcWebPort)
	if err := http.ListenAndServe(grpcWebPort, http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", corsOrigin)
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,x-grpc-web")
		wrappedGrpc.ServeHTTP(w, req)
	})); err != nil {
		log.Fatalf("failed to start grpc-web server: %v", err)
	}
}
