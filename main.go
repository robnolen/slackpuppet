package main

import (
	//"fmt"
	//"html/template"
	"log"
	"net/http"
	"os"
	//"path"
)

func handler(w http.ResponseWriter, r *http.Request) {

}

func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/", http.FileServer(http.Dir("html")))
	log.Fatal(http.ListenAndServe(":"+os.Getenv("VCAP_APP_PORT"), nil))
}
