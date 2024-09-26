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
