package main

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/lancer2672/DandelionServer_Go/api/server"
	"github.com/lancer2672/DandelionServer_Go/middleware"
	"github.com/lancer2672/DandelionServer_Go/pb/service"
	"github.com/lancer2672/DandelionServer_Go/server/sgrpc"
	"github.com/lancer2672/DandelionServer_Go/utils"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	serverConfig, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	conn, err := sql.Open(serverConfig.DBDriver, serverConfig.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database", err)
	}
	// runDatabaseMigration(serverConfig.MigrationUrl, serverConfig.DBSource)
	// go runGinServer(serverConfig, conn)
	go runGrpcServer(serverConfig, conn)
	runGatewayServer(serverConfig, conn)
	// runGinServer(serverConfig, conn)

}

func runGrpcServer(config utils.Config, conn *sql.DB) {
	server := sgrpc.NewServer(config, conn)
	grpcServer := grpc.NewServer()
	service.RegisterMovieServiceServer(grpcServer, server)

	//allow clients to see avaiable grpc server ~ self document
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal("Cannot create listener GRPC", err)
	}
	log.Println("GRPC Server started", config.GRPCServerAddress)
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("Cannot start GRPC", err)
	}

}

func runGatewayServer(config utils.Config, conn *sql.DB) {
	server := sgrpc.NewServer(config, conn)

	grpcMux := runtime.NewServeMux(
		//disable to keep snake_case variable names
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions: protojson.MarshalOptions{
				UseProtoNames: true,
			},
			UnmarshalOptions: protojson.UnmarshalOptions{
				DiscardUnknown: true,
			},
		}),
	)

	ctx, cancel := context.WithCancel(context.Background())
	//prevent system doing unnecessary works
	defer cancel()
	err := service.RegisterMovieServiceHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal("Cannot create register server")
	}
	mux := http.NewServeMux()

	mux.Handle("/", middleware.CheckApiKey(grpcMux))
	listener, err := net.Listen("tcp", config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot create listener HTTP Gateway", err)
	}
	log.Println("HTTP gateway Server started", config.ServerAddress)
	err = http.Serve(listener, mux)
	if err != nil {
		log.Fatal("Cannot start HTTP Gateway", err)
	}

}

func runGinServer(config utils.Config, conn *sql.DB) {
	server := server.NewServer(config, conn)
	err := server.Start()
	if err != nil {
		log.Fatal("Cannot start server", err)
	}
}

func runDatabaseMigration(url string, dbSource string) {
	m, err := migrate.New(
		url,
		dbSource)
	if err != nil {
		log.Fatal("cannot create migrate instance ", err)
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to migrate up", err)
	}
	log.Println("run db migration successfully")

}

// package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"mime"
// 	"net/http"
// 	"os"
// 	"path/filepath"
// 	"strconv"
// 	"strings"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/template/html/v2"
// )

// func main() {
// 	engine := html.New("./templates", ".tpl")
// 	app := fiber.New(fiber.Config{
// 		Views: engine,
// 	})

// 	// Define a route for the home page
// 	app.Get("/", func(c *fiber.Ctx) error {
// 		return c.Render("index", nil)
// 	})

// 	// Define a route for streaming video
// 	app.Get("/stream", streamVideo)

// 	// Start the Fiber server
// 	if err := app.Listen(":3000"); err != nil {
// 		log.Fatalf("unable to start the application , %v", err)
// 	}
// }

// func streamVideo(ctx *fiber.Ctx) error {

// 	filePath := "ex.mp4"

// 	// Open the video file
// 	file, err := os.Open(filePath)
// 	if err != nil {
// 		log.Println("Error opening video file:", err)
// 		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
// 	}
// 	defer file.Close()

// 	// Get the file size
// 	fileInfo, err := file.Stat()
// 	if err != nil {
// 		log.Println("Error getting file information:", err)
// 		return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
// 	}

// 	// get the file mime informations
// 	mimeType := mime.TypeByExtension(filepath.Ext(filePath))

// 	// get file size
// 	fileSize := fileInfo.Size()

// 	// Get the range header from the request
// 	// Get the range header from the request
// 	rangeHeader := ctx.Get(fiber.HeaderRange)
// 	if rangeHeader != "" {
// 		var start, end int64

// 		ranges := strings.Split(rangeHeader, "=")
// 		if len(ranges) != 2 {
// 			log.Println("Invalid Range Header:", err)
// 			return ctx.Status(http.StatusInternalServerError).SendString("Internal Server Error")
// 		}

// 		byteRange := ranges[1]
// 		byteRanges := strings.Split(byteRange, "-")

// 		// get the start range
// 		start, err := strconv.ParseInt(byteRanges[0], 10, 64)
// 		if err != nil {
// 			log.Println("Error parsing start byte position:", err)
// 			return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
// 		}

// 		// Calculate the end range
// 		if len(byteRanges) > 1 && byteRanges[1] != "" {
// 			end, err = strconv.ParseInt(byteRanges[1], 10, 64)
// 			if err != nil {
// 				log.Println("Error parsing end byte position:", err)
// 				return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
// 			}
// 		} else {
// 			end = fileSize - 1
// 		}

// 		// Setting required response headers
// 		ctx.Set(fiber.HeaderContentRange, fmt.Sprintf("bytes %d-%d/%d", start, end, fileInfo.Size())) // Set the Content-Range header
// 		ctx.Set(fiber.HeaderContentLength, strconv.FormatInt(end-start+1, 10))                        // Set the Content-Length header for the range being served
// 		ctx.Set(fiber.HeaderContentType, mimeType)                                                    // Set the Content-Type
// 		ctx.Set(fiber.HeaderAcceptRanges, "bytes")                                                    // Set Accept-Ranges
// 		ctx.Status(fiber.StatusPartialContent)                                                        // Set the status code to 206 (Partial Content)

// 		// Seek to the start position
// 		_, seekErr := file.Seek(start, io.SeekStart)
// 		if seekErr != nil {
// 			log.Println("Error seeking to start position:", seekErr)
// 			return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
// 		}

// 		// Copy the specified range of bytes to the response
// 		_, copyErr := io.CopyN(ctx.Response().BodyWriter(), file, end-start+1)
// 		if copyErr != nil {
// 			log.Println("Error copying bytes to response:", copyErr)
// 			return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
// 		}

// 	} else {
// 		// If no Range header is present, serve the entire video
// 		ctx.Set("Content-Length", strconv.FormatInt(fileSize, 10))
// 		_, copyErr := io.Copy(ctx.Response().BodyWriter(), file)
// 		if copyErr != nil {
// 			log.Println("Error copying entire file to response:", copyErr)
// 			return ctx.Status(fiber.StatusInternalServerError).SendString("Internal Server Error")
// 		}
// 	}

// 	return nil
// }
