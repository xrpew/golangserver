package handlers

import (
	"encoding/json"
	"main/models"
	"main/repository"
	"main/server"
	"net/http"

	"github.com/segmentio/ksuid"
)

type SingUpRequests struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SingUpResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func SingUpHandler(s server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requests = SingUpRequests{}
		err := json.NewDecoder(r.Body).Decode(&requests)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id, err := ksuid.NewRandom()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var user = models.User{
			Email:    requests.Email,
			Password: requests.Password,
			ID:       id.String(),
		}
		err = repository.InsertUser(r.Context(), &user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(SingUpResponse{ID: user.ID, Email: user.Email})

	}
}
