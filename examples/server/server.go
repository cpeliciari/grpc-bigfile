package main

import (
	"crypto/rand"
	"github.com/cpeliciari/grpc-bigfile"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

func main() {
	if err := generateFakeFile("file.bin", 10<<20); err != nil {
		log.Fatalf("Failed to generate fake file: %v", err)
	}

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	cfg := &bigfile.ServerConfig{
		Root:      "./",
		ChunkSize: 1 << 20, // 1MB
	}

	bigfile.StartServer(grpcServer, cfg)

	log.Println("Server is running on :50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func generateFakeFile(path string, size int32) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	data := make([]byte, size)
	rand.Read(data)
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	log.Printf("Fake file generated: %s\n", path)
	return nil
}
