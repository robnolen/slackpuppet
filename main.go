package main

import (
	"fmt"

	"bytes"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	type Page struct {
		Title string
		Body  []byte
	}
	p := Page{Title: "foo", Body: nil}
	r.ParseForm()
	message := []byte(r.FormValue("Message"))
	api_url := "https://devopsgeekweek.slack.com/services/hooks/slackbot?token=" + os.Getenv("SLACK_API_KEY")
	channel := "%23general"
	log.Println(api_url)
	resp, _ := http.Post(api_url+"&channel="+channel, "text/html", bytes.NewBuffer(message))
	response, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(response))
	t, _ := template.ParseFiles("html/index.html")
	t.Execute(w, p)
}
func main() {
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.Handle("/", http.FileServer(http.Dir("html")))
	http.HandleFunc("/sendmsg", handler)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("VCAP_APP_PORT"), nil))
}
