package main

import (
	"net/http"
)

type Login struct {
	HashedPassword string
	SessionToken   string
	CSRFToken      string
}

//key is the user name
var users = map[string]Login{}

//creating sample endpoints
func main() {
	http.HandleFunc("/register", register)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/protected", protected)
	http.ListenAndServe(":8080", nil)
}

//functions handling each request
func register(w http.ResponseWriter, r *http.Request) {
	//handles post requests
	if r.Method != http.MethodPost {
		er := http.StatusMethodNotAllowed
		http.Error(w, "Invalid Method", er)
		return
	}

	if err := r.ParseForm(); err != nil {
		er := http.StatusBadRequest
		http.Error(w, "Invalid form data", er)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")
	if len(password) < 8 {
		er := http.StatusNotAcceptable
		http.Error(w, "Invalid username / password", er)
		return
	}

	// checking if the user already exists
	if _, ok := users[username]; ok {
		er := http.StatusConflict
		http.Error(w, "user already exists", er)
		return
	}

	hashedPassword, err := hashPassword(password)
	if err != nil {
		er := http.StatusInternalServerError
		http.Error(w, "Failed to hash password", er)
		return
	}

	users[username] = Login{
		HashedPassword: hashedPassword,
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("registered"))
}
//logging a user in , session and crsf token used here
func login(w http.ResponseWriter, r *http.Request)     {




}
func logout(w http.ResponseWriter, r *http.Request)    {}
func protected(w http.ResponseWriter, r *http.Request) {}
