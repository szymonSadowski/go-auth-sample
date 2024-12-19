package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/goschool/crud/types"
)

func HandleEchoUser(w http.ResponseWriter, r *http.Request) {
	var user types.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "this is the email: %s", user.Email)
}

func HandleTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Test is working")
}
