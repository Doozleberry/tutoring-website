package main

import (
	"testing"
	"net/http"
	"fmt"
	"time"
)

func TestService(t *testing.T) {
	//go Run()
	time.Sleep(1000)
	resp, err := http.Get("http://localhost:8080")
	fmt.Println(resp)
	
	if err != nil {
		panic(err)
	}
	t.Fail()
}

