package main

import (
	"log"
	"net"

	"github.com/boris-wtf/secret-server/internal/notes"
	"github.com/boris-wtf/secret-server/pkg/github.com/boris-wtf/apis/secret"
	"github.com/spf13/viper"

	"github.com/boris-wtf/secret-server/internal"

	"google.golang.org/grpc"
)

func main() {
	err := initConfig()
	if err != nil {
		log.Fatal("failed to init config")
	}

	server := grpc.NewServer()
	secretService := &internal.SecretServiceServer{
		Notes: notes.NewRAMRegistry(viper.GetUint64("max_note_length")),
	}
	secret.RegisterSecretServiceServer(server, secretService)

	listener, err := net.Listen("tcp", "0.0.0.0:80")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Fatal(server.Serve(listener))
}
