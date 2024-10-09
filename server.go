package bigfile

import (
	"fmt"
	bigfilev1 "github.com/cpeliciari/grpc-bigfile/gen/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io"
	"os"
	"path/filepath"
)

const (
	defaultChunkSize = 1 << 20 // 1MB
	defaultMaxSend   = 10
	defaultRootDir   = "./"
)

type ServerConfig struct {
	Root      string
	ChunkSize int64
}

func (cfg *ServerConfig) defaults() {
	if cfg.ChunkSize == 0 {
		cfg.ChunkSize = defaultChunkSize
	}

	if cfg.Root == "" {
		cfg.Root = defaultRootDir
	}
}

func StartServer(grpcServer *grpc.Server, cfg *ServerConfig) {
	cfg.defaults()

	s := &server{
		chunkSize: cfg.ChunkSize,
		rootDir:   cfg.Root,
	}

	bigfilev1.RegisterFileTransferServiceServer(grpcServer, s)
	return
}

type server struct {
	chunkSize int64

	rootDir string
}

func (s *server) GetFile(request *bigfilev1.FileRequest, stream bigfilev1.FileTransferService_GetFileServer) error {
	filePath := filepath.Join(s.rootDir, request.GetFileName())

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		return fmt.Errorf("failed to get file stats: %v", err)
	}

	stat := &bigfilev1.FileResponse_Stat{
		LastModified: timestamppb.New(fileInfo.ModTime()),
		TotalSize:    int32(fileInfo.Size()),
	}

	totalChunks := int32(fileInfo.Size()/s.chunkSize) + 1

	if request.ChunkNumber >= totalChunks {
		return fmt.Errorf("requested chunk number %d is out of range, total chunks: %d", request.ChunkNumber, totalChunks)
	}

	offset := int64(request.ChunkNumber) * s.chunkSize
	_, err = file.Seek(offset, io.SeekStart)
	if err != nil {
		return fmt.Errorf("failed to seek to chunk %d: %v", request.ChunkNumber, err)
	}

	buffer := make([]byte, s.chunkSize)
	chunkNumber := request.ChunkNumber

	for {
		chunkNumber++
		n, err := file.Read(buffer)
		if err != nil && err != io.EOF {
			return fmt.Errorf("failed to read file: %v", err)
		}

		chunk := &bigfilev1.FileResponse_Chunk{
			ChunkNumber: chunkNumber,
			ChunkSize:   int32(n),
			TotalChunks: totalChunks,
			End:         chunkNumber == totalChunks,
			Data:        buffer[:n],
		}

		err = stream.Send(&bigfilev1.FileResponse{
			Name:  request.GetFileName(),
			Stat:  stat,
			Chunk: chunk,
		})
		if err != nil {
			return fmt.Errorf("failed to send chunk: %v", err)
		}

		if n == 0 {
			break
		}

	}

	return nil
}
