package main

import (
	"log"
	"net/http"
)

type Service struct {
	h func(http.ResponseWriter, *http.Request) error
}

func (h Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.h(w, r)
	if err != nil {
		log.Println("error:", err)
	}
}

func GetMux() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/", Service{GetResponse})
	return mux
}

func GetResponse(w http.ResponseWriter, r *http.Request) error {

	return nil
}

func main() {

}
