package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func tokenHandler(rw http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		badRequest(rw, "Only Post Allowed")
		return
	}

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		badRequest(rw, err.Error())
		return
	}

	bodyValues, err := url.ParseQuery(string(body))

	if err != nil {
		badRequest(rw, err.Error())
		return
	}

	if bodyValues["grant_type"] == nil {
		badRequest(rw, "Missing Grant Type")
		return
	}

	grantType := bodyValues["grant_type"][0]

	switch grantType {
	case "client_credentials":
		handleClientCredentials(rw, bodyValues)
	}

	_, _ = rw.Write([]byte("hello"))
}

func handleClientCredentials(rw http.ResponseWriter, bodyValues map[string][]string) {

}

func badRequest(rw http.ResponseWriter, message string) {
	http.Error(rw, message, http.StatusBadRequest)
}