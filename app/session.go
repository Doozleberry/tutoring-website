package main
//
//import (
//	"fmt"
//	"github.com/gorilla/mux"
//	"net/http"
//	"github.com/securecookie"
//)
//
//var cookieHandler = securecookie.New(
//	securecookie.GenerateRandomKey(64),
//	securecookie.GenerateRandomKey(32))
//
//func SetSession(userName string, response http.ResponseWriter) {
//	value := map[string]string{
//		"si": userName,
//	}
//	if encoded, err := cookieHandler.Encode("session", value); err == nil {
//		cookie := &http.Cookie{
//			Name:  "session",
//			Value: encoded,
//			Path:  "/",
//		}
//		http.SetCookie(response, cookie)
//	}
//}
//
//func ClearSession(response http.ResponseWriter) {
//	cookie := &http.Cookie{
//		Name:   "session",
//		Value:  "",
//		Path:   "/",
//		MaxAge: -1,
//	}
//	http.SetCookie(response, cookie)
//}

