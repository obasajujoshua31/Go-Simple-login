package handlers

import (
	"Go-Simple-login/providers"
	"Go-Simple-login/services"
	"fmt"
	"github.com/gorilla/sessions"
	"net/http"
)

const (
	password = "password"
	name = "name"
	email = "email"
	registerTemplate = "register"
	emailNotAvailable = "Email is not Available!"
	CookieName = "logged-in-status"
)

func GetRegisterUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		registerTemplate := ParseTemplate(registerTemplate)
		RenderTemplate(w, registerTemplate, nil)
   }
}


func PostRegisterUserHandler(connString string, store *sessions.CookieStore)  http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		email := GetInput(r, email)

		db, err := services.ConnectToDB(connString)

		if err != nil {
			RenderErrorPage(w, err)
			return
		}
		defer db.Close()

		user := db.FindUserByEmail(email)
		if user.Email != "" {
			registerTemplate := ParseTemplate(registerTemplate)
			RenderTemplate(w, registerTemplate, emailNotAvailable)
			return
		}

		password := GetInput(r,password)
		name := GetInput(r, name)
       newUser := services.User{
       	  Name: name,
       	  Password: password,
       	  Email: email,
	   }

	  err = db.CreateNewUser(newUser)
	  if err != nil {
		  RenderErrorPage(w, err)
		  return
	  }
	    providers.SetUserLoggedIn(w, r, store, CookieName)
	    http.Redirect(w, r, "/dashboard", http.StatusFound)
		return
	}
}

func GetInput(r *http.Request, name string) string{
	err := r.ParseForm()
	if err!= nil {
		fmt.Printf("Error %s", err)
	}
	return r.Form[name][0]
}

func RenderErrorPage(w http.ResponseWriter, err error) {
	errorTemplate := ParseTemplate("error")
	RenderTemplate(w, errorTemplate, err.Error())
}