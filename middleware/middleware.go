package middleware

import (
	"context"
	"fmt"

	apicalls "github.com/lancer2672/DandelionServer_Go/api_calls"
	"github.com/lancer2672/DandelionServer_Go/utils"
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
