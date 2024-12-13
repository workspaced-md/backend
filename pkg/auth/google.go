package auth

import (
	"context"
	"log"
	"net/http"
	"os"
	"encoding/json"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"github.com/labstack/echo/v4"
	"github.com/arnavsurve/workspaced/pkg/shared"
	"github.com/arnavsurve/workspaced/pkg/db"
)

var googleOauthConfig *oauth2.Config

// init api token from .env
func InitGoogleAuth() {
	googleOauthConfig = &oauth2.Config{
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URL"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.profile",
			"https://www.googleapis.com/auth/userinfo.email",
		},
		Endpoint: google.Endpoint,
	}
}

// redirects to login page
func HandleGoogleLogin(c echo.Context) error {
	url := googleOauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

// called after logging in. Creates/stores google user in db
func HandleGoogleCallback(c echo.Context, store *db.Store) error {
	// extracting code from http request param
	code := c.QueryParam("code")
	if code == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Authorization code not provided"})
	}

	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		log.Println("Failed to exchange token:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to login"})
	}

	client := googleOauthConfig.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Println("Failed to fetch user info:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to fetch user info"})
	}
	defer resp.Body.Close()

	user := struct {
		ID    string `json:"id"`
		Email string `json:"email"`
		Name  string `json:"name"`
	}{}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		log.Println("Failed to decode user info:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to decode user info"})
	}

	account, err := store.GetAccountByEmail(user.Email)
	if err != nil {
		// creates new user 
		account = &shared.Account{
			Email:    user.Email,
			Username: user.Name, 
		}

		if err := store.CreateAccount(account); err != nil {
			log.Println("Error saving user:", err)
			return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error saving user"})
		}
	}

	// Generate JWT
	tokenString, err := GenerateJWT(account)

	if err != nil {
		log.Println("Error generating token:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Error generating token"})
	}
	frontendURL := os.Getenv("FRONTEND_URL")
	redirectURL := frontendURL + "?token=" + tokenString

	return c.Redirect(http.StatusFound, redirectURL)
}
