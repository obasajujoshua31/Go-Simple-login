package providers

import (
    "github.com/gorilla/sessions"
    "net/http"
)

const authenticated = "authenticated"

func NewCookieStore(key string) *sessions.CookieStore {
     byteKey := []byte(key)
     return sessions.NewCookieStore(byteKey)
}

func SetUserLoggedIn(w http.ResponseWriter, r *http.Request, store *sessions.CookieStore, cookieName string) {
     session, _ := store.Get(r, cookieName)
     session.Values[authenticated] = true
     session.Save(r, w)
}

func LogOutUser(w http.ResponseWriter, r *http.Request, store *sessions.CookieStore, cookieName string) {
      session, _ := store.Get(r, cookieName)
      session.Values[authenticated] = false
     session.Save(r, w)
}

func IsUserLoggedIn(r *http.Request, store *sessions.CookieStore, cookieName string) bool{
    session, _ := store.Get(r, cookieName)
  auth, ok := session.Values[authenticated].(bool)


  if  !ok || !auth {
       return false
  }
  return true
}