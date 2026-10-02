package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/markbates/goth/gothic"
)

func (a *Application) createUserHandler(w http.ResponseWriter, r *http.Request) {

}


func (a *Application) getAuthCallback(w http.ResponseWriter, r *http.Request){
	provider := chi.URLParam(r, "provider")

	r = r.WithContext(context.WithValue(context.Background(), "provider", provider))

	user, err := gothic.CompleteUserAuth(w, r)
	if err != nil{
		fmt.Println(w, r)
		return
	}

	fmt.Println(user)

	http.Redirect(w, r, "idkdidk", http.StatusFound)
}