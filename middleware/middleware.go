package middleware

import (
	"context"
	"fmt"
	"time"

	"net/http"

	apicalls "github.com/lancer2672/DandelionServer_Go/api_calls"
	"github.com/lancer2672/DandelionServer_Go/utils"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "missing context metadata")
	}

	if len(md["x-api-key"]) != 1 {
		return nil, status.Errorf(codes.Unauthenticated, "invalid api key")
	}

	if md["x-api-key"][0] != "your-api-key" {
		return nil, status.Errorf(codes.Unauthenticated, "invalid api key")
	}

	return handler(ctx, req)
}

func CheckApiKey(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//check permission to add movie
		if r.Method == "POST" && r.URL.Path == "/movies" {
			apiKey := r.Header.Get("x-api-key")
			fmt.Println("APIHEADER", apiKey)
			role, permission, err := apicalls.GetInstance().CheckApiKey(apiKey)
			if err != nil {
				http.Error(w, "Invalid API Key", http.StatusUnauthorized)
				return
			}
			// 			editor_api_key
			// admin_api_key
			// user_api_key
			if role.Name == "admin_api_key" && utils.StringInSlice("movie", permission.Write) {
				fmt.Println("Permission granted")
			} else {
				http.Error(w, "Permission denied", http.StatusUnauthorized)
				return
			}
			fmt.Println("ROLE", role)
			fmt.Println("PERMISSION", permission)
		}

		next.ServeHTTP(w, r)
	})
}

func Logger(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		defer func() {
			log.Info().Str("method", r.Method).Str("path", r.URL.Path).Dur("duration", time.Since(start)).Msg("request")
		}()

		handler.ServeHTTP(w, r)
	})
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
		logger = logger.Err(err)
	}
	logger.Str("METHOD", info.FullMethod).Int("STATUS", int(statusCode)).Str("CODE", statusCode.String()).Dur("DURATION", duration).Msg("")
	return result, err
}
