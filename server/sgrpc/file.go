package sgrpc

import (
	"io"
	"log"
	"os"

	"github.com/lancer2672/DandelionServer_Go/pb/service"

	"github.com/lancer2672/DandelionServer_Go/pb/request"
)

func (server *Server) SendFile(r *request.SendFileRequest, stream service.MovieService_SendFileServer) error {
	file, err := os.OpenFile("ex.mp4", os.O_RDWR, 0666)
	if err != nil {
		log.Fatal("Cannot open file", err)
	}
	defer file.Close()
	if err != nil {
		log.Fatal("Cannot get file info", err)
	}

	buffer := make([]byte, 1024)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		feature := &request.SendFileResponse{
			Chunk: buffer[:n],
		}
		if err := stream.Send(feature); err != nil {
			return err
		}
	}
	return nil
}
