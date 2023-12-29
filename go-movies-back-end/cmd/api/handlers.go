package main

import (
	"errors"
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
	var requestPayload struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := app.readJSON(w, r, &requestPayload)
	if err != nil {
		app.errorJSON(w, err, http.StatusBadRequest)
		return
	}

	// Validate user against DB
	user, err := app.DB.GetUserByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("invalid email id"), http.StatusBadRequest)
		return
	}
	log.Println("***** handlers-authenticate-user: ", user)

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
	log.Println("***** handlers-authenticate-tokens: ", tokens)
	log.Println("***** handlers-authenticate-AccessToken: ", tokens.Token)
	log.Println("***** handlers-authenticate-RefreshToken: ", tokens.RefreshToken)

	refreshCookie := app.auth.GetRefreshCookie(tokens.RefreshToken)
	log.Println("***** handlers-authenticate-refreshCookie: ", refreshCookie)

	http.SetCookie(w, refreshCookie)
	w.Write([]byte(tokens.Token))
}
