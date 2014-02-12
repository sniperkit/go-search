package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/codegangsta/martini"
)

var (
	fileFlag = flag.String("file", "shakespeare.txt", "file to read")
	portFlag = flag.String("port", "8080", "port to run the server on")
)

func setupConfig() {
	cmd := os.Args[0]
	flag.Usage = func() {
		fmt.Println(`Usage:`, cmd, `[<options>] 
Options:`)
		flag.PrintDefaults()
	}
	flag.Parse()
}

func main() {
	setupConfig()
	m := setupMartini()
	http.ListenAndServe(":"+*portFlag, m)
}

func setupMartini() (m *martini.ClassicMartini) {
	m = martini.Classic()
	m.Use(SearchEngine())
	m.Get("/search.json", SearchHandler)
	m.Get("/", MainHandler)
	return
}
