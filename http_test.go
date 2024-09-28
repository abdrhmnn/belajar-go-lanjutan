package belajargolanjutan

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
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

// query parameter
func UrlParameter(respon http.ResponseWriter, request *http.Request) {
	nama := request.URL.Query().Get("nama")
	umur := request.URL.Query().Get("umur")

	// get multiple value parameter
	query := request.URL.Query()
	key := query["hobby"]

	fmt.Fprintf(respon, "Hello, %s! umur aku %s \n", nama, umur)
	fmt.Fprint(respon, "Hobby aku: ")
	fmt.Fprint(respon, strings.Join(key, ", "))
}

func TestQueryParameter(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?nama=abdu&umur=100&hobby=game&hobby=mancing", nil)
	recorder := httptest.NewRecorder()

	UrlParameter(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
}

// header
func HandlerHeader(respon http.ResponseWriter, request *http.Request) {
	header := request.Header.Get("content-type")

	// kirim header dari server
	respon.Header().Add("duarrrr", "eunha")

	fmt.Fprint(respon, header)
}

func TestHeader(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080", nil)

	// kirim header dari client
	request.Header.Add("Content-Type", "abdu")
	recorder := httptest.NewRecorder()

	HandlerHeader(recorder, request)

	response := recorder.Result()
	fmt.Println(response.Header.Get("duarrrr"))
}

// post form
func PostData(respon http.ResponseWriter, request *http.Request) {
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	nama := request.PostForm.Get("nama")
	umur := request.PostForm.Get("umur")

	fmt.Fprintf(respon, "Hello, %s, umur %s", nama, umur)
}

func TestPostData(t *testing.T) {
	requestBody := strings.NewReader("nama=abdu&umur=100")
	requestUrl := httptest.NewRequest(http.MethodPost, "http://localhost:8080", requestBody)
	requestUrl.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	PostData(recorder, requestUrl)

	respon := recorder.Result()
	body, _ := io.ReadAll(respon.Body)

	fmt.Println(string(body))
}

// handling response code
func ResponCode(respon http.ResponseWriter, request *http.Request) {
	nama := request.URL.Query().Get("nama")
	if nama == "" {
		respon.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(respon, "not found!")
	} else {
		respon.WriteHeader(http.StatusOK)
		fmt.Fprintf(respon, "Hello, %s", nama)
	}
}

func TestResponseCode(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080?nama=abdu", nil)
	recorder := httptest.NewRecorder()

	ResponCode(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))
	fmt.Println(response.StatusCode)
}
