package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("The port is not found in the environment")
	}

	r := chi.NewRouter()

	// middleware
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("helo"))
	})

	v1R := chi.NewRouter()
	v1R.Get("/healthz", handlerFunction)
	v1R.Get("/err", handlerErr)
	r.Mount("/v1", v1R)

	server := &http.Server{
		Handler: r,
		Addr:    ":" + portString,
	}
	fmt.Printf("Listening on port: %s", portString)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
