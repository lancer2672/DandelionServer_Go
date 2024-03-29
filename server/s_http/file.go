package shttp

import (
	"fmt"
	"io"
	"net/http"
)

func StreamFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Lấy movieUrl từ tham số truy vấn
	movieUrl := r.URL.Query().Get("movieUrl")
	if movieUrl == "" {
		http.Error(w, "Movie URL is required", http.StatusBadRequest)
		return
	}
	fmt.Println("MovieUrl", movieUrl)
	resp, err := http.Get(movieUrl)
	if err != nil {
		http.Error(w, "Cannot open URL", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	if _, err := io.Copy(w, resp.Body); err != nil {
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}
