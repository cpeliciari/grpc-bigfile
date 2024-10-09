package bigfile

import (
	"context"
	"fmt"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"time"

	bigfilev1 "github.com/cpeliciari/grpc-bigfile/gen/proto/v1"
	"google.golang.org/grpc"
)

type ClientInterface interface {
	GetFile(ctx context.Context, filename string, chunkNumber int32) (<-chan Response, error)
	Close() error
}

type base struct {
	client bigfilev1.FileTransferServiceClient
	conn   *grpc.ClientConn
}

type ClientConfig struct {
	ServerAddress string
}

func NewClient(cfg *ClientConfig) (ClientInterface, error) {
	conn, err := grpc.NewClient(cfg.ServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("falha ao conectar: %v", err)
	}
	client := bigfilev1.NewFileTransferServiceClient(conn)
	return &base{client: client, conn: conn}, nil
}

func (b *base) Close() error {
	return b.conn.Close()
}

type Response struct {
	Name         string
	LastModified time.Time
	Size         int64
	ChunkNumber  int32
	ChunkSize    int32
	TotalChunks  int32
	End          bool
	Data         []byte
}

func (b *base) GetFile(ctx context.Context, filename string, chunkNumber int32) (<-chan Response, error) {
	request := &bigfilev1.FileRequest{
		FileName:    filename,
		ChunkNumber: chunkNumber,
	}

	stream, err := b.client.GetFile(ctx, request)
	if err != nil {
		return nil, fmt.Errorf("failed to get file: %v", err)
	}

	responseChan := make(chan Response)

	go func() {
		defer close(responseChan)
		for {
			response, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				log.Printf("error receiving chunk: %v", err)
				return
			}

			chunk := response.GetChunk()
			stat := response.GetStat()
			resp := Response{
				Name:         filename,
				LastModified: stat.GetLastModified().AsTime(),
				Size:         int64(stat.GetTotalSize()),
				ChunkNumber:  chunk.GetChunkNumber(),
				ChunkSize:    chunk.GetChunkSize(),
				TotalChunks:  chunk.GetTotalChunks(),
				End:          chunk.GetEnd(),
				Data:         chunk.GetData(),
			}
			responseChan <- resp
		}
	}()

	return responseChan, nil
}
