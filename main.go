package main

import (
	"fmt"
	"net/http"
	"sync"
)

var validUserIDs sync.Map

func main() {
	http.HandleFunc("/dyn-user", HelloUser)
	http.HandleFunc("/check-userid", CheckUserID)
	http.HandleFunc("/add-user", AddUser)
	http.HandleFunc("/delete-user", DeleteUser)
	http.ListenAndServe("0.0.0.0:8080", http.DefaultServeMux)
}

func HelloUser(resp http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("userid")
	fmt.Fprintf(resp, "Hello, %s\n", id)
}

func CheckUserID(resp http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("userid")
	if isUserIDValid(id) {
		fmt.Fprintln(resp, "Hello, user!")
	} else {
		http.Error(resp, "Bad user ID", http.StatusForbidden)
	}
}

func AddUser(resp http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id != "" {
		validUserIDs.Store(id, true)
		fmt.Fprintf(resp, "User %s added successfully!\n", id)
	} else {
		http.Error(resp, "Invalid user ID", http.StatusBadRequest)
	}
}

func DeleteUser(resp http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("id")
	if id != "" {
		validUserIDs.Delete(id)
		fmt.Fprintf(resp, "User %s deleted successfully!\n", id)
	} else {
		http.Error(resp, "Invalid user ID", http.StatusBadRequest)
	}
}

func isUserIDValid(id string) bool {
	_, ok := validUserIDs.Load(id)
	return ok
}
