package main

import (
	"log"
	"noteify-api/database"
	"noteify-api/routers"
)

func main() {
	db, err := database.ConnectDB()
	log.Print(db)
	if err != nil {
		panic("failed to connect database")
	}
	r := routers.SetupRouter()

	// Route for initiating the OAuth2 flow
	r.GET("/auth/google/login", handleGoogleLogin)

	// Route for handling the callback from Google
	r.GET("/auth/google/callback", handleGoogleCallback)

	r.Run()
}
