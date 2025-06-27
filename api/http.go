package api

import (
	"encoding/json"
	"log"
	"net/http"
	"webapp/domain"

	_ "github.com/go-chi/chi/v5/middleware"
)

type UserStorageIF interface {
	// NewUserStorage() *gorm.DB
	//All() ([]storage.UserDBModel, error)
	All() (*[]domain.User, error)
}

type UserHandler struct {
	storage UserStorageIF
}

func NewUserHandler(storage UserStorageIF) *UserHandler {
	return &UserHandler{
		storage: storage,
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	usersPtr, err := h.storage.All()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var usersResponse []UserResponse
	users := *usersPtr
	for _, u := range users {
		usersResponse = append(usersResponse, userResponseFromDomainModel(u))
	}

	err = json.NewEncoder(w).Encode(usersResponse)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func userResponseFromDomainModel(u domain.User) UserResponse {
	var emails []EmailResponse
	for _, e := range u.Emails {
		emails = append(emails, emailResponseFromDomainModel(e))
	}

	return UserResponse{
		Id:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Emails:    emails,
	}
}

func emailResponseFromDomainModel(e domain.Email) EmailResponse {
	return EmailResponse{
		Address: e.Address,
		Primary: e.Primary,
	}
}
