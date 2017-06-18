package main

import (
	//"fmt"
	"html/template"
	"net/http"
)

var homeView = template.Must(template.ParseFiles("exhibit-d/home.html"))
var template_vars TemplateVars

type TemplateVars struct {
	nameList map[string]int
	testy    string
}

func home(resp http.ResponseWriter, req *http.Request) {
	homeView.Execute(resp, &template_vars)
}

func addUser(name string) {
	if template_vars.nameList[name] >= 1 {
		template_vars.nameList[name]++
	} else {
		template_vars.nameList = make(map[string]int)
		template_vars.nameList[name] = 1
	}
	return
}

func signup(resp http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	username := req.Form.Get("username")

	addUser(username)
	http.Redirect(resp, req, "/home", http.StatusFound)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/home", home)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8080", nil)
}
