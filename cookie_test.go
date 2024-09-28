package belajargolanjutan

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(respon http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "Testing"
	cookie.Value = "Abdu"
	cookie.Path = "/"

	http.SetCookie(respon, cookie)
	fmt.Fprint(respon, "Cookie berhasil dibuat!")
}

func GetCookie(respon http.ResponseWriter, request *http.Request) {
	cookie, _ := request.Cookie("Testing")
	fmt.Fprintf(respon, "Hello, %s", cookie.Value)
}

func TestLiveCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

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

func TestCookieServer(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	recorder := httptest.NewRecorder()

	// set cookie di server dan mengirim ke client
	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookie %s %s", cookie.Name, cookie.Value)
	}
}

func TestCookieClient(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)
	cookie := new(http.Cookie)
	cookie.Name = "Testing"
	cookie.Value = "Eunha"
	cookie.Path = "/"

	// set cookie di client dan mengirim ke server
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}
