package main

import (
	"fmt"
	_ "github.com/lib/pq"
	"github.com/tutoring-website/util"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

var valid = false

func signupHandler(w http.ResponseWriter, r *http.Request) {
	var (
		firstname = r.FormValue("firstName")
		lastname  = r.FormValue("lastName")
		email     = r.FormValue("email")
		username  = r.FormValue("userName")
		password1 = r.FormValue("pw1")
		//password2 = r.FormValue("pw2")
	)

	r.ParseForm()
	signupPassword := []byte(password1)
	hashedSignupPassword, check := bcrypt.GenerateFromPassword(signupPassword, bcrypt.DefaultCost)
	util.CheckErr(check)
	fmt.Println(string(hashedSignupPassword))
	check = bcrypt.CompareHashAndPassword(hashedSignupPassword, signupPassword)
	fmt.Println(check) // nil means it is a match

	var lastInsertId string
	err := db.QueryRow("INSERT INTO users(first_name, last_name, email, username, password) VALUES($1,$2,$3,$4,$5) returning username;", firstname, lastname, email, username, hashedSignupPassword).Scan(&lastInsertId)
	util.CheckErr(err)
	fmt.Println("last inserted id =", lastInsertId)
	//w.WriteHeader(http.StatusCreated)
}

func signinHandler(w http.ResponseWriter, r *http.Request) {
	var dbpassword []byte
	username := r.FormValue("suUsername")
	password := r.FormValue("suPassword")
	signinPassword := []byte(password)
	rows, err := db.Query("select username, password from users where username = $1;", username)
	util.CheckErr(err)
	log.Println(err)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&username, &dbpassword)
		util.CheckErr(err)
		log.Println(username)
		log.Println(dbpassword)
		log.Println(signinPassword)
		util.CheckErr(bcrypt.CompareHashAndPassword(dbpassword, signinPassword)) //if no message about nonmatch, successful login verification achieved
		//util.CheckErr(bcrypt.CompareHashAndPassword(dbpassword, hashedSigninPassword))
		if err == nil {
		valid = true
		}
	}
	err = rows.Err()
	util.CheckErr(err)
	fmt.Println(valid)
	if valid == true {
		http.Redirect(w, r, "http://localhost:8080/homepage/v.html", 302)
}
}

func Redirect (w http.ResponseWriter, r *http.Request) {
	
	http.Redirect(w, r, "http://www.golang.org", 301)
}


//func signoutHandler(response http.ResponseWriter, request *http.Request) {
//	clearSession(response)
//	http.Redirect(response, request, "/", 302)
//}

//func homePageHandler(response http.ResponseWriter, request *http.Request) {
//	userName := getUserName(request)
//	fmt.Println(valid)
//	if valid == true {
//		fmt.Fprintf(response, internalPage, userName)
//	} else {
//		http.Redirect(response, request, "/", 302)
//	}
//}