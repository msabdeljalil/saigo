package main

import (
	//"fmt"
	"html/template"
	"net/http"
)

type TemplateVars struct {
	nameList map[string]int
}

var homeView = template.Must(template.ParseFiles("exhibit-d/home.html"))
var template_vars TemplateVars

func home(resp http.ResponseWriter, req *http.Request) {
	homeView.Execute(resp, &template_vars)
}

func signup(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	username := req.Form.Get("username")

	addUser(username)
	http.Redirect(resp, req, "/home", http.StatusFound)
}

func addUser(name string) {
	template_vars.nameList[name]++
}

func main() {
	template_vars.nameList = make(map[string]int)

	http.HandleFunc("/", home)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8080", nil)
}
