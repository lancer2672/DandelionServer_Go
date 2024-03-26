package middleware

import (
	"context"
	"fmt"
	"time"

	apicalls "github.com/lancer2672/DandelionServer_Go/api_calls"
	"github.com/lancer2672/DandelionServer_Go/utils"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func CheckApiKeyInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "missing context metadata")
	}
	if len(md["x-api-key"]) != 1 {
		return nil, status.Errorf(codes.Unauthenticated, "invalid api key")
	}
	apiKey := md["x-api-key"][0]
	fmt.Println("GRPC Api Header", apiKey)
	fmt.Println("GRPC md", md)
	_, permission, err := apicalls.CheckApiKey(apiKey)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Invalid API Key")
	}

	if apiKey == "admin_api_key" && utils.StringInSlice("movie", permission.Write) {
		log.Error().Err(err).Msg("Permission granted")
	} else {
		return nil, status.Errorf(codes.Unauthenticated, "Permission denied")
	}

	return handler(ctx, req)
}

func LoggerInterceptor(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	start := time.Now()
	result, err := handler(ctx, req)
	duration := time.Since(start)
	statusCode := codes.Unknown
	if st, ok := status.FromError(err); ok {
		statusCode = st.Code()
	}
	logger := log.Info()
	if err != nil {
		logger = log.Error().Err(err)
	}
	logger.Str("METHOD", info.FullMethod).Int("STATUS", int(statusCode)).Str("CODE", statusCode.String()).Dur("DURATION", duration).Msg("")
	return result, err
}
