package main

import (
	"context"
	bigfile "github.com/cpeliciari/grpc-bigfile"
	"log"
	"time"
)

func main() {
	cfg := &bigfile.ClientConfig{
		ServerAddress: "localhost:50051",
	}

	client, err := bigfile.NewClient(cfg)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()

	// Request file starting from chunk 0
	responseChan, err := client.GetFile(ctx, "file.bin", 0)
	if err != nil {
		log.Fatalf("Failed to get file: %v", err)
	}

	for resp := range responseChan {
		log.Printf("Received chunk: %d, Total Chunks: %d, Chunk Size: %d", resp.ChunkNumber, resp.TotalChunks, resp.ChunkSize)
	}
}
