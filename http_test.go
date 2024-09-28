package belajargolanjutan

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HandlerMethod(respon http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(respon, "Hello World!")
}

func TestHttpTesting(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	HandlerMethod(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
