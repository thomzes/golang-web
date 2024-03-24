package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {

	var handler http.HandlerFunc = func (writer http.ResponseWriter, request *http.Request)  {
		// logic web
		fmt.Fprint(writer, "Hello World")
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}

func TestServeMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// logic web
		fmt.Fprint(w, "Hello World!")
	})

	mux.HandleFunc("/hi", func(w http.ResponseWriter, r *http.Request) {
		// logic web
		fmt.Fprint(w, "Hi")
	})

	mux.HandleFunc("/images/", func(w http.ResponseWriter, r *http.Request) {
		// logic web
		fmt.Fprint(w, "IMAGES!")
	})

	mux.HandleFunc("/images/thumbnails/", func(w http.ResponseWriter, r *http.Request) {
		// logic web
		fmt.Fprint(w, "THUMBNAILS!")
	})
	
	server := http.Server {
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestRequest(t *testing.T) {
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, r.Method)
		fmt.Fprint(w, r.RequestURI)
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}


