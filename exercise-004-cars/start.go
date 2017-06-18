package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sync"
)

// Vehicle ...
type Vehicle struct {
	Name  string
	Count int
}

// Vehicles ...
type Vehicles struct {
	List []*Vehicle
}

// View ...
type View struct {
	Username string
	Vehicles Vehicles
}

// Globals
var joinT *template.Template
var playT *template.Template
var sessions map[string]*Vehicles
var lock sync.Mutex

// Setup (Globals)
func setup(dir string) {
	joinT = template.Must(template.ParseFiles(dir + "/templates/join.html"))
	playT = template.Must(template.ParseFiles(dir + "/templates/play.html"))
	sessions = make(map[string]*Vehicles)
}

// Action: Home
func home(w http.ResponseWriter, r *http.Request) {
	joinT.Execute(w, nil)
}

// Action: Join
func join(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	cookie := http.Cookie{Name: "username", Value: username}
	http.SetCookie(w, &cookie)

	// Protect your criticals!
	lock.Lock()
	sessions[username] = &Vehicles{}
	lock.Unlock()

	http.Redirect(w, r, "/play", http.StatusFound)
}

// Action: Play
func play(w http.ResponseWriter, r *http.Request) {
	username, _ := r.Cookie("username")

	lock.Lock()
	vehicles, _ := sessions[username.Value]
	lock.Unlock()

	view := View{Username: username.Value, Vehicles: *vehicles}
	playT.Execute(w, &view)
}

// Add ...
func (v *Vehicles) Add(name string) {
	for _, vehicle := range v.List {
		if vehicle.Name == name {
			vehicle.Count++
			return
		}
	}
	vehicle := &Vehicle{Name: name, Count: 1}
	v.List = append(v.List, vehicle)
}

// Action: Add
func add(w http.ResponseWriter, r *http.Request) {
	username, _ := r.Cookie("username")

	r.ParseForm()
	vehicle := r.Form.Get("vehicle")
	speed := r.Form.Get("vehicle")
	name := vehicle + "_" + speed

	vehicles, _ := sessions[username.Value]
	vehicles.Add(name)

	view := View{Username: username.Value}
	playT.Execute(w, &view)
}

// Action: Exit
func exit(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/", http.StatusFound)
}

func main() {
	setup(".")

	// Static Files
	staticHandler := http.StripPrefix("/public/", http.FileServer(http.Dir("./public")))
	http.Handle("/public/", staticHandler)

	http.HandleFunc("/", home)
	http.HandleFunc("/join", join)
	http.HandleFunc("/play", play)
	http.HandleFunc("/add", add)
	http.HandleFunc("/exit", exit)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
	}
}
