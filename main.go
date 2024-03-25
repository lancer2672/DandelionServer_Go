package main

import (
	"context"
	"database/sql"

	"net"
	"net/http"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/lancer2672/DandelionServer_Go/middleware"
	"github.com/lancer2672/DandelionServer_Go/pb/service"
	"github.com/lancer2672/DandelionServer_Go/server/sgrpc"
	"github.com/lancer2672/DandelionServer_Go/utils"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
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
		log.Error().Err(err).Msg("Error loading config")
	}
	conn, err := sql.Open(serverConfig.DBDriver, serverConfig.DBSource)
	if err != nil {
		log.Error().Err(err).Msg("Cannot connect to database")
		os.Exit(1)
	}
	// runDatabaseMigration(serverConfig.MigrationUrl, serverConfig.DBSource)
	// go runGinServer(serverConfig, conn)
	go runGrpcServer(serverConfig, conn)
	runGatewayServer(serverConfig, conn)
	// runGinServer(serverConfig, conn)

}

func runGrpcServer(config utils.Config, conn *sql.DB) {
	server := sgrpc.NewServer(config, conn)
	loggerInterceptor := grpc.UnaryInterceptor(middleware.LoggerInterceptor)
	//allow clients to see avaiable grpc server ~ self document
	grpcServer := grpc.NewServer(loggerInterceptor)
	service.RegisterMovieServiceServer(grpcServer, server)
	reflection.Register(grpcServer)
	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Error().Err(err).Msg("Cannot create listener GRPC")
		os.Exit(1)
	}
	log.Info().Str("address", config.GRPCServerAddress).Msg("GRPC Server started")
	if err = grpcServer.Serve(listener); err != nil {
		log.Error().Err(err).Msg("Cannot start GRPC")
		os.Exit(1)
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
		log.Error().Err(err).Msg("Cannot create listener HTTP Gateway")
		os.Exit(1)
	}

	mux := http.NewServeMux()

	mux.Handle("/", middleware.Logger(middleware.CheckApiKey(grpcMux)))
	listener, err := net.Listen("tcp", config.ServerAddress)
	if err != nil {
		log.Error().Err(err).Msg("Cannot create listener HTTP Gateway")
	}
	// err = http.Serve(listener, mux)
	log.Info().Str("address", config.ServerAddress).Msg("HTTP gateway Server started")
	if err = http.Serve(listener, mux); err != nil {
		log.Error().Err(err).Msg("Cannot start HTTP Gateway")
		os.Exit(1)
	}

}

func runDatabaseMigration(url string, dbSource string) {
	m, err := migrate.New(
		url,
		dbSource)
	if err != nil {
		log.Error().Err(err).Msg("Cannot create migrate instance")
		os.Exit(1)
	}
	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Error().Err(err).Msg("Failed to migrate up")
		os.Exit(1)
	}
	log.Info().Msg("Run db migration successfully")

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
