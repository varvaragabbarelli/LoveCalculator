package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"

	"github.com/varvaragabbarelli/LoveCalculator/websiteContent"
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
	details := websiteContent.WebContent{
		Name1: r.FormValue("fname"),
		Name2: r.FormValue("sname"),
	}

	log.Println("name1: ", details.Name1, " Name2: ", details.Name2)

	resp, _ := websiteContent.DoRequest(&details)

	body, _ := ioutil.ReadAll(resp.Body)
	details, _ = websiteContent.GetResult(string(body))
	tmpl.Execute(w, struct {
		Success bool
		Result  string
	}{true, details.Result})
	return nil
}

func main() {
	http.ListenAndServe(fmt.Sprintf(":%d", 8080), GetMux())
}
