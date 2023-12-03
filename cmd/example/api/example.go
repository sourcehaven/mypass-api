package main

import (
	"github.com/rs/zerolog/log"
)

func main() {
	baseurl := "http://127.0.0.1:7277"
	apiurl := baseurl + "/api"
	authurl := apiurl + "/user"
	_ = apiurl + "/vault"

	registerBody := []byte(`{"username":"User","password":"password"}`)
	registerResp := MakePost(authurl+"/register", registerBody)
	log.Print(registerResp)

	loginBody := []byte(`{username: "User", "password": "password"}`)
	loginResp := MakePost(authurl+"/login", loginBody)

	log.Print(loginResp)
}
