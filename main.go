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

// Middleware to validate the user query parameter
// ValidateUser function here is added as middleware. It checks if the "userid" query
// parameter is valid. If the parameter is valid, it calls the next handler function.
// So this way user IDs are validated before the actual endpoint logic is executed. (All handlers are included into the middleware)
func ValidateUser(next http.HandlerFunc) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		id := req.URL.Query().Get("userid")
		if id == "" || !isUserIDValid(id) {
			http.Error(resp, "Invalid user ID", http.StatusForbidden)
			return
		}
		next(resp, req)
	}
}

func HelloUser(resp http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("userid")
	fmt.Fprintf(resp, "Hello, %s\n", id)
}

func CheckUserID(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(resp, "Hello, user!")
}

func AddUser(resp http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("userid")
	validUserIDs.Store(id, true)
	fmt.Fprintf(resp, "User %s added successfully!\n", id)
}

func DeleteUser(resp http.ResponseWriter, req *http.Request) {
	id := req.URL.Query().Get("userid")
	validUserIDs.Delete(id)
	fmt.Fprintf(resp, "User %s deleted successfully!\n", id)
}

func isUserIDValid(id string) bool {
	_, ok := validUserIDs.Load(id)
	return ok
}
