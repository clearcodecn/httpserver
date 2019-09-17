package main

import (
	"flag"
	"net/http"
	"os"
)

var (
	addr string
	staticDir string
)

func init() {
	flag.StringVar(&addr,"addr",":8080","address")
	flag.StringVar(&staticDir,"s","/static","static dir")
}
func main() {
	flag.Parse()
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		name, _ := os.Hostname()
		writer.Write([]byte(name))
	})

	http.Handle("/static",http.FileServer(http.Dir(staticDir)))


	http.ListenAndServe(addr, nil)
}
