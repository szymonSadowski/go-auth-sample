package api

import (
	"encoding/json"
	"net/http"

	"github.com/goschool/crud/db"
	"github.com/goschool/crud/types"
)

type UserHandler struct {
	userStore db.UserStore
}

type LoginParams struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	User  *types.User `json:"user"`
	Token string      `json:"token"`
}

func NewUserHandler(us db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: us,
	}
}

func (u *UserHandler) HandlerRegisterUser(w http.ResponseWriter, r *http.Request) {
	var createUser types.CreateUser
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&createUser); err != nil {
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		return
	}

	user, err := types.NewUser(createUser)

	if err != nil {
		http.Error(w, "Invalid new user", http.StatusBadRequest)
		return
	}
	newUser, err := u.userStore.CreateUser(ctx, user)
	if err != nil {
		http.Error(w, "Invalid new user", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(newUser); err != nil {
		http.Error(w, "could not register new user", http.StatusInternalServerError)
	}
}

func (u *UserHandler) HandlerLoginUser(w http.ResponseWriter, r *http.Request) {
	var params LoginParams
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		http.Error(w, "Something went wrong", http.StatusBadRequest)
		return
	}
	// does this user actually exisits in our db
	user, err := u.userStore.GetUserByEmail(ctx, params.Email)
	if err != nil {
		http.Error(w, "User not found", http.StatusInternalServerError)
		return
	}
	// return the relevant info
	// validate password
	if !types.ValidatePassword(user.PasswordHash, params.Password) {
		http.Error(w, "Invalid credentials", http.StatusInternalServerError)
		return
	}

	response := LoginResponse{
		User:  user,
		Token: types.CreateToken(*user),
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "could not register new user", http.StatusInternalServerError)
	}

}
