package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
)

// ============================================================================================
func (app *application) Home(w http.ResponseWriter, r *http.Request) {

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

// ============================================================================================
func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {

	log.Println("***** handlers-AllMovies app: ", app)
	//log.Println("***** handlers-AllMovies-request: ", r)
	//log.Println("***** handlers-AllMovies-request header: ", r.Header)
	//log.Println("***** handlers-AllMovies-request body: ", r.Body)
	//log.Println("***** handlers-AllMovies-*request: ", *r)
	//log.Println("***** handlers-AllMovies-r.cookies: ", r.Cookies())

	for _, cookie := range r.Cookies() {
		log.Println("***** handlers-AllMovies-cookie: ", cookie)
		log.Println("***** handlers-AllMovies-cookie name: ", cookie.Name)
	}

	movies, err := app.DB.AllMovies()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, movies) //this is new addition
}

// ============================================================================================
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

	// Generate the refresh cookie
	refreshCookie := app.auth.GetRefreshCookie(tokens.RefreshToken)
	log.Println("***** handlers-authenticate-refreshCookie: ", refreshCookie)

	// Create the cookie
	http.SetCookie(w, refreshCookie)
	app.writeJSON(w, http.StatusAccepted, tokens)
}

// ============================================================================================
func (app *application) refreshToken(w http.ResponseWriter, r *http.Request) {

	log.Println("***** handlers-refreshToken-r.cookies: ", r.Cookies())
	log.Println("***** handlers-refreshToken-request: ", r.Body)

	for _, cookie := range r.Cookies() {
		if cookie.Name == app.auth.CookieName {
			claims := &Claims{}
			refreshToken := cookie.Value

			// parse the token to get the claims
			// ParseWithClaims(refreshToken, claims,  (interface{}, error) {
			_, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(app.JWTSecret), nil
			})
			if err != nil {
				app.errorJSON(w, errors.New("unauthorized"), http.StatusUnauthorized)
				return
			}

			// get the user id from the token claims
			userID, err := strconv.Atoi(claims.Subject)
			if err != nil {
				app.errorJSON(w, errors.New("unknown user"), http.StatusUnauthorized)
				return
			}

			user, err := app.DB.GetUserByID(userID)
			if err != nil {
				app.errorJSON(w, errors.New("unknown user"), http.StatusUnauthorized)
				return
			}

			u := jwtUser{
				ID:        user.ID,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			}

			tokenPairs, err := app.auth.GenerateTokenPair(&u)
			if err != nil {
				app.errorJSON(w, errors.New("error generating tokens"), http.StatusUnauthorized)
				return
			}

			http.SetCookie(w, app.auth.GetRefreshCookie(tokenPairs.RefreshToken))

			app.writeJSON(w, http.StatusOK, tokenPairs)

		}
	}

}

// ============================================================================================
func (app *application) logout(w http.ResponseWriter, r *http.Request) {

	log.Println("***** handlers-logout-r.Cookies(): ", r.Cookies())
	http.SetCookie(w, app.auth.GetExpiredRefreshCookie())
	w.WriteHeader(http.StatusAccepted)

}

// ============================================================================================
func (app *application) MovieCatalog(w http.ResponseWriter, r *http.Request) {
	movies, err := app.DB.AllMovies()
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	_ = app.writeJSON(w, http.StatusOK, movies)
}
