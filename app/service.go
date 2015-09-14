package main

import (
"net/http"
	
)


func Run() error {
	defer db.Close()
	router := NewRouter()
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./ui/")))
	//router.HandleFunc("/signin", signinHandler)
	http.Handle("/", router)
	return http.ListenAndServe(":8080", nil)
}