package main

import (
	"fmt"

	//"html/template"
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	//"path"
)

func handler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	message := []byte(r.FormValue("Message"))
	api_url := "https://devopsgeekweek.slack.com/services/hooks/slackbot?token=ADywv0e2hv25DM3ea18VylIa"
	channel := "%23general"
	log.Println(api_url)
	resp, _ := http.Post(api_url+"&channel="+channel, "text/html", bytes.NewBuffer(message))
	response, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(response))
}
func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/", http.FileServer(http.Dir("html")))
	http.HandleFunc("/sendmsg", handler)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("VCAP_APP_PORT"), nil))
}
