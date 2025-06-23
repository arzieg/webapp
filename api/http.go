package api

import (
	"net/http"
	"webapp/storage"

	"github.com/go-chi/chi/v5"
	_ "github.com/go-chi/chi/v5/middleware"
)

type ServerInterface interface {
	// Get all users
	// (GET /users)
	GetUsers(w http.ResponseWriter, r *http.Request)
	// Add a new user
	// (POST /users)
	// PostUser(w http.ResponseWriter, r *http.Request)
	// Delete user
	// (DELETE /users/{userID})
	// DeleteUser(w http.ResponseWriter, r *http.Request, userID UserID)
	// Get a single user
	// (GET /users/{userID})
	// GetUser(w http.ResponseWriter, r *http.Request, userID UserID)
	// Update user
	// (PATCH /users/{userID})
	// PatchUser(w http.ResponseWriter, r *http.Request, userID UserID)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
}

type MiddlewareFunc func(http.HandlerFunc) http.HandlerFunc

func (siw *ServerInterfaceWrapper) GetUsers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var handler = func(w http.ResponseWriter, r *http.Request) {
		siw.Handler.GetUsers(w, r)
	}

	for _, middleware := range siw.HandlerMiddlewares {
		handler = middleware(handler)
	}

	handler(w, r.WithContext(ctx))
}

type UserHandler struct {
	storage storage.UserStorage
}

func NewUserHandler(storage storage.UserStorage) UserHandler {
	return UserHandler{
		storage: storage,
	}
}

// Handler creates http.Handler with routing matching OpenAPI spec.
func Handler(si ServerInterface) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{})
}

type ChiServerOptions struct {
	BaseURL     string
	BaseRouter  chi.Router
	Middlewares []MiddlewareFunc
}

// HandlerFromMux creates http.Handler with routing matching OpenAPI spec based on the provided mux.
func HandlerFromMux(si ServerInterface, r chi.Router) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseRouter: r,
	})
}

func HandlerFromMuxWithBaseURL(si ServerInterface, r chi.Router, baseURL string) http.Handler {
	return HandlerWithOptions(si, ChiServerOptions{
		BaseURL:    baseURL,
		BaseRouter: r,
	})
}

// HandlerWithOptions creates http.Handler with additional options
func HandlerWithOptions(si ServerInterface, options ChiServerOptions) http.Handler {
	r := options.BaseRouter

	if r == nil {
		r = chi.NewRouter()
	}
	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
	}

	r.Group(func(r chi.Router) {
		r.Get(options.BaseURL+"/users", wrapper.GetUsers)
	})
	/*
		r.Group(func(r chi.Router) {
			r.Post(options.BaseURL+"/users", wrapper.PostUser)
		})
		r.Group(func(r chi.Router) {
			r.Delete(options.BaseURL+"/users/{userID}", wrapper.DeleteUser)
		})
		r.Group(func(r chi.Router) {
			r.Get(options.BaseURL+"/users/{userID}", wrapper.GetUser)
		})
		r.Group(func(r chi.Router) {
			r.Patch(options.BaseURL+"/users/{userID}", wrapper.PatchUser)
		})
	*/

	return r
}
