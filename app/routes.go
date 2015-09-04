package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
//	Route{
//		"indexHome",
//		"GET",
//		"/",
//		indexHandler,
//	},
	Route{
		"saveSignup",
		"POST",
		"/signup",
		signupHandler,
	},
	Route{
		"matchSignin",
		"POST",
		"/signin",
		signinHandler,
	},
//	Route{
//		"clearSession",
//		"/logout",
//		"POST",
//		signoutHandler,
//	},
//	Route{
//		"homePage",
//		"/homepage",
//		homePageHandler,
//	},
}
