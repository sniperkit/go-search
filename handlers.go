package main

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/argusdusty/Ferret"
	"github.com/codegangsta/martini"
)

var templates = template.Must(template.ParseFiles("main.tmpl"))

func MainHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "main.tmpl", nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func SearchHandler(res http.ResponseWriter, log *log.Logger, req *http.Request, params martini.Params, se *ferret.InvertedSuffix) {
	t := time.Now()
	query := req.FormValue("term")
	results, values := se.Query(query, 10)
	timing := time.Now().Sub(t).String()
	log.Printf("Searched for %v in %v\n", query, timing)
	data, _ := json.Marshal(SearchResponse{query, timing, results, values})
	res.Write(data)
}
