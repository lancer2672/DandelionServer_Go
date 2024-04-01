package shttp

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func StreamFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	movieUrl := r.URL.Query().Get("movieUrl")
	if movieUrl == "" {
		http.Error(w, "Movie URL is required", http.StatusBadRequest)
		return
	}
	fmt.Println("MovieUrl", movieUrl)
	// resp, err := http.Get(movieUrl)
	// if err != nil {
	// 	http.Error(w, "Cannot open URL", http.StatusInternalServerError)
	// 	return
	// }
	// defer resp.Body.Close()

	file, err := os.OpenFile("ex.mp4", os.O_RDONLY, 0666)
	if err != nil {
		log.Fatal("Cannot open file", err)
	}
	defer file.Close()
	fmt.Println("reading")
	if _, err := io.Copy(w, file); err != nil {
		fmt.Println("err", err)
		http.Error(w, "Error writing response", http.StatusInternalServerError)
		return
	}
}
