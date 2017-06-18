package main

import (
	//"fmt"
	"html/template"
	"net/http"
)

type View struct {
	nameList map[string]int
	testy    string
}

var homeT = template.Must(template.ParseFiles("exhibit-d/home.html"))
var v View

func home(w http.ResponseWriter, r *http.Request) {
	homeT.Execute(w, &v)
}

func addUser(name string) {
	if v.nameList[name] >= 1 {
		v.nameList[name]++
	} else {
		v.nameList = make(map[string]int)
		v.nameList[name] = 1
	}
	return
}

func signup(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")

	addUser(username)
	http.Redirect(w, r, "/home", http.StatusFound)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/home", home)
	http.HandleFunc("/signup", signup)
	http.ListenAndServe(":8080", nil)
}
