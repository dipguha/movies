package main

import (
	"log"
	"net/http"
)

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello World from %s", app.Domain)
	var payload = struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Version string `json:"version"`
	}{
		Status:  "Active",
		Message: "Go movies up and running",
		Version: "1.0.0",
	}

	_ = app.writeJSON(w, http.StatusOK, payload)

}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	log.Println("***** handlers - AllMovies app *****: ", app)
	movies, err := app.DB.AllMovies()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, movies) //this is new addition
}

func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	// Read the json payload

	// Validate user against DB

	// Check password

	// Create a JWT user
	u := jwtUser{
		ID:        1,
		FirstName: "Admin",
		LastName:  "User",
	}

	// Generate token
	tokens, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	log.Println("***** handlers - authenticate tokens *****: ", tokens)
	log.Println("***** handlers - authenticate Access token *****: ", tokens.Token)
	log.Println("***** handlers - authenticate Refresh token *****: ", tokens.RefreshToken)

	w.Write([]byte(tokens.Token))
}
