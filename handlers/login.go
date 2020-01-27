package handlers

import (
	"Go-Simple-login/providers"
	"Go-Simple-login/services"
	"fmt"
	"github.com/gorilla/sessions"
	"html/template"
	"net/http"
)

const (
	loginTemplate = "login"
	invalidEmail = "Invalid Login Credentials!"
)
func GetLoginHandler(store *sessions.CookieStore)  http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if status := providers.IsUserLoggedIn(r, store, CookieName); status {
			http.Redirect(w, r, "/dashboard", http.StatusFound)
			return
		}

		loginTemplate := ParseTemplate("login")
		RenderTemplate(w, loginTemplate, nil)
	}
}

func PostLoginHandler(connString string, store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		email := GetInput(r, "email")

		db, err := services.ConnectToDB(connString)
		if err != nil {
			RenderErrorPage(w, err)
			return
		}

		defer db.Close()

		user := db.FindUserByEmail(email)
		if user.Email == "" {
			registerTemplate := ParseTemplate(loginTemplate)
			RenderTemplate(w, registerTemplate, invalidEmail)
			return
		}

		if auth:= user.IsMatchPassword(GetInput(r, "password")); !auth {
			registerTemplate := ParseTemplate(loginTemplate)
			RenderTemplate(w, registerTemplate, invalidEmail)
			return
		}

		providers.SetUserLoggedIn(w, r, store, CookieName)
		http.Redirect(w, r, "/dashboard", http.StatusFound)
		return
	}
}

func LogoutHandler(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		providers.LogOutUser(w, r, store, CookieName)
		http.Redirect(w, r, "/login", http.StatusFound)
		return
	}
}

func RenderTemplate (w http.ResponseWriter, t *template.Template, data interface{}){
		if err:= t.Execute(w, data); err != nil {
			http.Error(w, fmt.Sprintf("Error executing template %s", err), http.StatusInternalServerError)
		}
}

func ParseTemplate(page string) *template.Template {
	return template.Must(template.New(fmt.Sprintf("%s.gohtml", page)).ParseGlob("templates/*.gohtml"))
}
