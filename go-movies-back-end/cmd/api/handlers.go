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

// ------------------------------------------------------------------------------------
func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {

	log.Println("***** handlers-authenticate-r: ", r)

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

	log.Println("***** handlers-authenticate-requestPayload: ", requestPayload)

	// Validate user against DB
	user, err := app.DB.GetUserByEmail(requestPayload.Email)
	if err != nil {
		app.errorJSON(w, errors.New("invalid email id"), http.StatusBadRequest)
		return
	}
	log.Println("***** handlers-authenticate-user: ", user)

	// Check password
	valid, err := user.PasswordMatches(requestPayload.Password)

	if err != nil || !valid {
		app.errorJSON(w, errors.New("passwords don't match"), http.StatusBadRequest)
		return
	}

	// Create a JWT user
	u := jwtUser{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
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
	app.writeJSON(w, http.StatusAccepted, tokens)
}

// ------------------------------------------------------------------------------------
func (app *application) refreshToken(w http.ResponseWriter, r *http.Request) {

	log.Println("***** handlers-refreshToken-start r.cookies: ", r.Cookies())

	// Loop through all cookies
	for _, cookie := range r.Cookies() {
		log.Println("***** handlers-refreshToken-cookie: ", cookie)
		log.Println("***** handlers-refreshToken-cookie name: ", cookie.Name)
	}

}

// ------------------------------------------------------------------------------------
func (app *application) logout(w http.ResponseWriter, r *http.Request) {
	log.Println("***** handlers-logout-r.Cookies(): ", r.Cookies())
	http.SetCookie(w, app.auth.GetExpiredRefreshCookie())
	w.WriteHeader(http.StatusAccepted)

}
