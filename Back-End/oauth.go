package main

import (
	"context"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"net/http"
	"noteify-api/repositories"
	"os"
)

var (
	googleOauthConfig *oauth2.Config
	oauthStateString  = "pseudo-random"
)

type UserProfile struct {
	Email          string `json:"email"`
	Name           string `json:"name"`
	ProfilePicture string `json:"picture"`
}

func init() {
	err := godotenv.Load()
	if err != nil {
		panic("failed to load env file")
	}
	clientid := os.Getenv("GOOGLE_CLIENT_ID")
	clientsecret := os.Getenv("GOOGLE_CLIENT_SECRET")
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8080/auth/google/callback",
		ClientID:     clientid,
		ClientSecret: clientsecret,
		Scopes: []string{
			"openid",
			"profile",
			"email",
		},
		Endpoint: google.Endpoint,
	}
}

func handleGoogleLogin(c *gin.Context) {
	url := googleOauthConfig.AuthCodeURL(oauthStateString)
	c.Redirect(http.StatusTemporaryRedirect, url)
}

func handleGoogleCallback(c *gin.Context) {
	state := c.Query("state")
	if state != oauthStateString {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	code := c.Query("code")
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// Fetch the user's profile information from the Google API
	client := googleOauthConfig.Client(context.Background(), token)
	userInfo, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer userInfo.Body.Close()

	var userProfile UserProfile

	err = json.NewDecoder(userInfo.Body).Decode(&userProfile)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	// add user to database
	repositories.CreateUserFromParams(userProfile.Email, userProfile.Name)

	// generate jwt
	jwt, err := generateJWT(userProfile)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": jwt})
}

func generateJWT(userProfile UserProfile) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": userProfile.Email,
		"name":  userProfile.Name,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
