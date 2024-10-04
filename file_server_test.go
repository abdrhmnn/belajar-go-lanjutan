package belajargolanjutan

import (
	"embed"
	"fmt"
	"io/fs"
	"net/http"
	"testing"
)

func TestFileServer(t *testing.T) {
	directory := http.Dir("./resources")
	fileServer := http.FileServer(directory)

	// bisa menggunakan Serve biasa atau ServeMux
	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	fmt.Println("server is running!")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// file server with golang embed

//go:embed resources
var resources embed.FS

func TestFileServerWithEmbed(t *testing.T) {
	directory, _ := fs.Sub(resources, "resources") // untuk masuk lagi ke dalam folder resources
	fileServer := http.FileServer(http.FS(directory))

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}

	fmt.Println("server is running!")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
