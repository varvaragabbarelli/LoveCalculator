package main

import (
	"LoveCalculator/web"
	"log"
	"net/http"
	"text/template"
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
	tmpl := template.Must(template.ParseFiles("tmp/ui.html"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return nil
	}
	details := web.WebResult{
		Name1: r.FormValue("fname"),
		Name2: r.FormValue("sname"),
	}
	return nil
}

func main() {

}
