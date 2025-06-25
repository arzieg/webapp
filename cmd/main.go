package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"webapp/api"
	"webapp/models"
	"webapp/storage"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = models.AutoMigrate(db)
	if err != nil {
		log.Fatalf("failed to apply migration. Got %v", err)
	}

	storage := storage.NewUserStorage(db)

	h := api.NewUserHandler(storage)

	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", h.GetUsers)
		/*		r.Post("/", h.PostUser)

				r.Route("/{userID}", func(r chi.Router) {
					r.Get("/", h.GetUser)
					r.Patch("/", h.PatchUser)
					r.Delete("/", h.DeleteUser)
				})
		*/
	})

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
