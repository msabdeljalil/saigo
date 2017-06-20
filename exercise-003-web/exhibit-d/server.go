package main

import (
	"html/template"
	"net/http"
)

var homeView = template.Must(template.ParseFiles("exhibit-d/home.html"))
var NameList map[string]int = make(map[string]int)

func home(resp http.ResponseWriter, req *http.Request) {
	for k, v := range NameList {
		println(k, v)
	}
	homeView.Execute(resp, NameList)
}

func signup(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	username := req.Form.Get("username")

	addUser(username)
	http.Redirect(resp, req, "/home", http.StatusFound)
}

func addUser(name string) {
	NameList[name]++
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8080", nil)
}
