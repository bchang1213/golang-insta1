package main

import (
	"dep-ex/internal/routes"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: false,
		AllowOriginFunc:  func(origin string) bool { return true },
		Debug:            true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
	})
	// http.Handle("/", routes.Index())
	handler := c.Handler(routes.Index())
	var port = "8080"
	err := http.ListenAndServe("localhost:"+port, handler)
	if err != nil {
		panic(err)
		return
	}
}
