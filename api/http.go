package api

import (
	"encoding/json"
	"log"
	"net/http"
	"webapp/models"
	"webapp/storage"

	_ "github.com/go-chi/chi/v5/middleware"
)

type UserHandler struct {
	storage storage.UserStorage
}

func NewUserHandler(storage storage.UserStorage) UserHandler {
	return UserHandler{
		storage: storage,
	}
}

func (h UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.storage.All()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var usersResponse []UserResponse
	for _, u := range users {
		usersResponse = append(usersResponse, userResponseFromDBModel(u))
	}

	err = json.NewEncoder(w).Encode(usersResponse)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func userResponseFromDBModel(u models.UserDBModel) UserResponse {
	var emails []EmailResponse
	for _, e := range u.Emails {
		emails = append(emails, emailResponseFromDBModel(e))
	}

	return UserResponse{
		Id:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Emails:    emails,
	}
}

func emailResponseFromDBModel(e models.EmailDBModel) EmailResponse {
	return EmailResponse{
		Address: e.Address,
		Primary: e.Primary,
	}
}
