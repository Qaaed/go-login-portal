package main

import (
	"fmt"
	"net/http"
	"time"
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
	if r.Method != http.MethodPost{
		er := http.StatusMethodNotAllowed
		http.Error(w,"invalid request method",er)
		return
	}
	username := r.FormValue("username")
	password := r.FormValue("password")

	user,ok := users[username]
	if !ok || !checkPasswordHash(password,user.HashedPassword){
		er:=http.StatusUnauthorized
		http.Error(w,"invalid username or password",er)
		return
	}

	fmt.Fprintln(w,"Login Successful!")

	sessionToken := generateToken(32)
	csrfToken := generateToken(32)

	//setting a session cookie
	//cookies work by sending a request from the browser to the backened everytime a request is made
	http.SetCookie(w,&http.Cookie{
		Name : "session_token",
		Value: sessionToken,
		Expires: time.Now().Add(24 * time.Hour), //any request after 24 hours will get auto logged out
		HttpOnly: true, //ensures the session token can't be accessed by frontend js in the client
	})
	//setting the csrf token inside acookie
		http.SetCookie(w,&http.Cookie{
		Name : "csrf_token",
		Value: csrfToken,
		Expires: time.Now().Add(24 * time.Hour), //any request after 24 hours will get auto logged out
		HttpOnly: false, //asking if should be accessible by client side
	})

	//storing the session token
	user.SessionToken = sessionToken
	user.CSRFToken = csrfToken
	users[username] = user
	fmt.Fprintln(w,"Login Successful!")



}
func logout(w http.ResponseWriter, r *http.Request)    {}
func protected(w http.ResponseWriter, r *http.Request) {}
