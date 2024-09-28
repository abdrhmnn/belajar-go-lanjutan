package belajargolanjutan

import (
	"fmt"
	"net/http"
	"testing"
)

func TestServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
	}

	fmt.Println("server is running!")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// handler respon dari request client
func TestHandlerRequest(t *testing.T) {
	// http handlerFunc adalah implementasi dari interface Handler yaitu root interface untuk handling
	// respon dan request dri client
	var handler http.HandlerFunc = func(respon http.ResponseWriter, request *http.Request) {
		fmt.Fprint(respon, "Hello World!")
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	fmt.Println("server is running!")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

// ServeMux / router
func TestServerMux(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Route home!")
	})
	mux.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Route about!")
	})
	mux.HandleFunc("/profile/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Route profile!")
	})
	mux.HandleFunc("/profile/detail/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Route detail profile!")
	})

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

// get request
func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(respon http.ResponseWriter, request *http.Request) {
		fmt.Fprint(respon, request.Method)
		fmt.Fprint(respon, request.RequestURI)
	}

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: handler,
	}

	fmt.Println("server is running!")
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
