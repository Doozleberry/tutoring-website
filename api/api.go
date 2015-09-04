package api

import (

)

type User struct {
	FirstName string `json :"firstName"`
	LastName string `json :"lastName"`
	Email string `json :"email"`
	Username string `json :"userName"`
	Password string `json :"pw1"`
}

