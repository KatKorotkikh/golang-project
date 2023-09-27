package main

import (
	"fmt"
	"net/http"
)

var validUserIDs = []string{
	"user1",
	"user2",
	"user3",
}

func main() {
	http.HandleFunc("/dyn-user", HelloUser)
	http.HandleFunc("/check-userid", CheckUserID)
	http.ListenAndServe("0.0.0.0:8080", http.DefaultServeMux)
}

func HelloUser(resp http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("userid")
	fmt.Fprintf(resp, "Hello, %s\n", id)
}

func CheckUserID(resp http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("userid")
	if isUserIdValid(id) {
		fmt.Fprintln(resp, "Hello, user!")
	} else {
		http.Error(resp, "Bad user ID", http.StatusForbidden)
	}
}

func isUserIdValid(id string) bool {
	for _, validUser := range validUserIDs {
		if id == validUser {
			return true
		}
	}
	return false
}
