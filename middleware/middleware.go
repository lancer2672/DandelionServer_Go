package middleware

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"net/http"
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
		apiKey := r.Header.Get("x-api-key")
		fmt.Println("APIHEADER", apiKey)
		// if apiKey != "admin_api_key" && apiKey != "editor_api_key" && apiKey != "user_api_key" {
		// 	http.Error(w, "Invalid API Key", http.StatusUnauthorized)
		// 	return
		// } else {
		// }
		next.ServeHTTP(w, r)
	})
}
