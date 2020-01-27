package handlers

import (
	"Go-Simple-login/providers"
	"github.com/gorilla/sessions"
	"net/http"
)

func HomeHandler(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		status := providers.IsUserLoggedIn(r, store, CookieName)

		if status {
			http.Redirect(w, r, "/dashboard", http.StatusFound)
			return
		}

		homeTemplate := ParseTemplate("home")
		RenderTemplate(w, homeTemplate, nil)
		return
	}
}

func GetDashboardHandler(store *sessions.CookieStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if status := providers.IsUserLoggedIn(r, store, CookieName); !status {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		dasTemplate := ParseTemplate("dashboard")
		RenderTemplate(w, dasTemplate, true )
	}
}

