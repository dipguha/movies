package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Auth struct {
	Issuer        string
	Audience      string
	Secret        string
	TokenExpiry   time.Duration
	RefreshExpiry time.Duration
	CookieDomain  string
	CookiePath    string
	CookieName    string
}

type jwtUser struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type TokenPairs struct {
	Token        string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type Claims struct {
	jwt.RegisteredClaims
}

func (j *Auth) GenerateTokenPair(user *jwtUser) (TokenPairs, error) {
	// Create an JWT access token
	token := jwt.New(jwt.SigningMethodHS256)
	log.Println("***** Auth-GenerateTokenPair-AccessToken: ***** ", token)

	// Set the claim
	claims := token.Claims.(jwt.MapClaims)
	claims["name"] = fmt.Sprintf("%s %s", user.FirstName, user.LastName)
	claims["sub"] = fmt.Sprint(user.ID)
	claims["aud"] = j.Audience
	claims["iss"] = j.Issuer
	claims["iat"] = time.Now().UTC().Unix()
	claims["typ"] = "JWT"

	// Set the expiry for JWT access token
	claims["exp"] = time.Now().UTC().Add(j.TokenExpiry).Unix()

	// Create a signed token
	signedAccessToken, err := token.SignedString([]byte(j.Secret))
	if err != nil {
		return TokenPairs{}, err
	}
	log.Println("***** Auth-GenerateTokenPair-signedAccessToken: ", signedAccessToken)

	// Create a refresh token and set claims
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	log.Println("***** Auth-GenerateTokenPair-RefreshToken: ", refreshToken)

	refreshTokenClaims := refreshToken.Claims.(jwt.MapClaims)
	refreshTokenClaims["sub"] = fmt.Sprint(user.ID)
	refreshTokenClaims["iat"] = time.Now().UTC().Unix()

	// Set the expiry of the refresh token
	refreshTokenClaims["exp"] = time.Now().UTC().Add(j.RefreshExpiry).Unix()

	// Create signed refresh token
	signedRefreshToken, err := refreshToken.SignedString([]byte(j.Secret))
	if err != nil {
		return TokenPairs{}, err
	}
	log.Println("***** Auth-GenerateTokenPair-signedRefreshToken: ", signedRefreshToken)

	// Create TokenPairs and populate with signed tokens
	var tPairs = TokenPairs{
		Token:        signedAccessToken,
		RefreshToken: signedRefreshToken,
	}

	// Return TokenPairs
	return tPairs, nil
}

func (j *Auth) GetRefreshCookie(refreshToken string) *http.Cookie {
	/* 	refreshCookie := &http.Cookie{
	   		Name:     j.CookieName,
	   		Path:     j.CookiePath,
	   		Value:    refreshToken,
	   		Expires:  time.Now().Add(j.RefreshExpiry),
	   		MaxAge:   int(j.RefreshExpiry.Seconds()),
	   		SameSite: http.SameSiteStrictMode,
	   		Domain:   j.CookieDomain,
	   		HttpOnly: true,
	   		Secure:   true,
	   	}
	   	log.Println("***** Auth-GetRefreshCookie-refreshCookie: ", refreshCookie)
	   	return refreshCookie */

	return &http.Cookie{
		Name:     j.CookieName,
		Path:     j.CookiePath,
		Value:    refreshToken,
		Expires:  time.Now().Add(j.RefreshExpiry),
		MaxAge:   int(j.RefreshExpiry.Seconds()),
		SameSite: http.SameSiteStrictMode,
		Domain:   j.CookieDomain,
		HttpOnly: true,
		Secure:   true,
	}
}

func (j *Auth) GetExpiredRefreshCookie(refreshToken string) *http.Cookie {
	return &http.Cookie{
		Name:     j.CookieName,
		Path:     j.CookiePath,
		Value:    "",
		Expires:  time.Unix(0, 0),
		MaxAge:   -1,
		SameSite: http.SameSiteStrictMode,
		Domain:   j.CookieDomain,
		HttpOnly: true,
		Secure:   true,
	}
	//log.Println("***** Auth-GetRefreshCookie-refreshCookie: ", refreshCookie)
}
