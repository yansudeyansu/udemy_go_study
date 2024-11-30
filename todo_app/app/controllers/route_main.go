package controllers

import (
	"net/http"
)

func top(w http.ResponseWriter, r *http.Request) {
	_, err := session(w, r)
	if err != nil {
		GenerateHTML(w, "Hello", "layout", "public_navbar", "top")
	} else {
		http.Redirect(w, r, "/todo", 302)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	sess, err := session(w, r)
	if err != nil {
		http.Redirect(w, r, "/", 302)
	} else {
		GenerateHTML(w, sess, "layout", "private_navbar", "index")
	}
}
