package auth

import (
	"github.com/gorilla/sessions"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	googleprovider "github.com/markbates/goth/providers/google"
	"golang.org/x/oauth2"
)

var(
	key = "randomstring"
	MaxAge = 86400 * 30
	IsProd = false
)

func NewAuth(cfg *oauth2.Config) {

	store := sessions.NewCookieStore([]byte(key))
	store.MaxAge(MaxAge)

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = IsProd

	gothic.Store = store

	goth.UseProviders(
		googleprovider.New(cfg.ClientID, cfg.ClientSecret, "idsksnd"),
	)
}