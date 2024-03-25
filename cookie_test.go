package golang_web

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCokie(writer http.ResponseWriter, request *http.Request) {
	cookie := new(http.Cookie)
	cookie.Name = "X-THOMAS-Name"
	cookie.Value = request.URL.Query().Get("name")
	cookie.Path = "/"

	http.SetCookie(writer, cookie)
	fmt.Fprintf(writer, "Success create cookie!")
}

func GetCookie(writer http.ResponseWriter, request *http.Request) {
	cookie, err := request.Cookie("X-THOMAS-Name")
	if err != nil {
		fmt.Fprintf(writer, "Cookie Not Found!")
	} else {
		fmt.Fprintf(writer, "Hello %s", cookie.Value)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCokie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}	
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http:/localhost:8080/?name=Thomas", nil)
	recoder := httptest.NewRecorder()

	SetCokie(recoder, request)

	cookies := recoder.Result().Cookies()

	for _, cookie := range cookies {
		fmt.Printf("Cookie %s : %s", cookie.Name, cookie.Value)
	}
}

func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest(http.MethodGet, "http:/localhost:8080", nil)

	cookie := new(http.Cookie)
	cookie.Name = "X-THOMAS-Name"
	cookie.Value = "Thomas"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}