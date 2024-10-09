# grpc-bigfile

[![Go Reference](https://pkg.go.dev/badge/github.com/cpeliciari/grpc-bigfile.svg)](https://pkg.go.dev/github.com/cpeliciari/grpc-bigfile)

`grpc-bigfile` is a gRPC-based service for transferring large files in chunks. It consists of a server that serves files and a client that requests files in chunks over gRPC.

## Features
- File transfer in chunks to efficiently handle large files
- Server configuration options for chunk size and root directory
- Client can request specific file chunks

## Installation

To use this package, you need to install Go and set up your environment.

```bash
go get github.com/cpeliciari/grpc-bigfile
```

## Usage

### Server

The server is responsible for serving files from a specified directory in chunks. By default, it uses a chunk size of 1MB and serves files from the current directory (`./`).

#### Configuration

- `Root`: The root directory from which the server will serve files (default: `./`).
- `ChunkSize`: The size of each chunk in bytes (default: `1MB`).

#### Start the Server

```go
package main

import (
	"log"
	"net"
	"github.com/cpeliciari/grpc-bigfile"
	"google.golang.org/grpc"
)

func main() {
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
```

### Client

The client requests a file from the server in chunks. The client receives the file chunks in a stream and can process them as needed.

#### Example

```go
package main

import (
	"context"
	"log"
	"time"
	"github.com/cpeliciari/grpc-bigfile"
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
```

### Protocol Definition (`proto`)

The file transfer protocol is defined in a gRPC `.proto` file. This defines the `FileTransferService` service and message formats for file requests and responses.

```proto
syntax = "proto3";

package bigfile.v1;

import "google/protobuf/timestamp.proto";

service FileTransferService {
  rpc GetFile(FileRequest) returns (stream FileResponse) {};
}

message FileRequest {
  string file_name = 1;
  int32 chunk_number = 2;
}

message FileResponse {
  message Stat {
    google.protobuf.Timestamp last_modified = 1;
    int32 total_size = 2;
  }

  message Chunk {
    int32 chunk_number = 1;
    int32 chunk_size = 2;
    int32 total_chunks = 3;
    bool end = 4;
    bytes data = 5;
  }

  string name = 1;
  Stat stat = 2;
  Chunk chunk = 3;
}
```

## Example of Use

Here’s a full example demonstrating how to generate a file, start the server, and retrieve the file via the client.

```go
package main

import (
	"context"
	"crypto/rand"
	"log"
	"net"
	"os"
	"time"
	"github.com/cpeliciari/grpc-bigfile"
	"google.golang.org/grpc"
)

const (
	serverAddress = "localhost:50051"
	fileName      = "file.bin"
	chunkSize     = 1 << 20  // 1MB
	fakeFileSize  = 10 << 20 // 10MB
)

func main() {
	if err := generateFakeFile(fileName); err != nil {
		log.Fatalf("Failed to generate fake file: %v", err)
	}

	lis, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	cfgServer := &bigfile.ServerConfig{
		Root:      "./",
		ChunkSize: chunkSize,
	}

	bigfile.StartServer(grpcServer, cfgServer)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	cfg := &bigfile.ClientConfig{
		ServerAddress: serverAddress,
	}

	client, err := bigfile.NewClient(cfg)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute*10)
	defer cancel()

	responseChan, err := client.GetFile(ctx, fileName, 0)
	if err != nil {
		log.Fatalf("Failed to get file: %v", err)
	}

	var totalReceived int64

	for resp := range responseChan {
		totalReceived += int64(len(resp.Data))
		log.Printf("Received chunk %+v\n", resp)
	}

	log.Printf("Total received: %d bytes\n", totalReceived)
}

func generateFakeFile(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	data := make([]byte, fakeFileSize)
	rand.Read(data)
	_, err = file.Write(data)
	if err != nil {
		return err
	}

	log.Printf("Fake file generated: %s\n", path)
	return nil
}
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
