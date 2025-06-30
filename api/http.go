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
	Add(user *domain.User) error
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

func userDomainModelFromCreateRequest(r domain.CreateUserRequest) domain.User {
	return domain.User{
		FirstName: r.FirstName,
		LastName:  r.LastName,
		Emails: []domain.Email{
			{
				Address: r.Email,
			},
		},
	}
}

func (h UserHandler) PostUser(w http.ResponseWriter, r *http.Request) {
	var createUserRequest domain.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&createUserRequest)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//validate := validator.New()
	// err = validate.Struct(createUserRequest)
	// if err != nil {
	// 	log.Println(err)
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	user := userDomainModelFromCreateRequest(createUserRequest)

	err = h.storage.Add(&user)

	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
