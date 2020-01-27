package main

import (
     "Go-Simple-login/config"
     "Go-Simple-login/db"
     "Go-Simple-login/handlers"
     "Go-Simple-login/providers"
     "fmt"
     "github.com/gorilla/mux"
     "github.com/gorilla/sessions"
     "log"
     "net/http"
)

type Router struct {
     *mux.Router
}

func newRouter() *Router {
     return &Router{mux.NewRouter()}
}

func (r *Router) setupRoutes(conString string, sessionStore *sessions.CookieStore) {
     r.HandleFunc("/", handlers.HomeHandler(sessionStore)).Methods(http.MethodGet)
     r.HandleFunc("/login", handlers.GetLoginHandler(sessionStore)).Methods(http.MethodGet)
     r.HandleFunc("/login", handlers.PostLoginHandler(conString, sessionStore)).Methods(http.MethodPost)
     r.HandleFunc("/register", handlers.GetRegisterUserHandler()).Methods(http.MethodGet)
     r.HandleFunc("/register", handlers.PostRegisterUserHandler(conString, sessionStore)).Methods(http.MethodPost)
     r.HandleFunc("/logout", handlers.LogoutHandler(sessionStore)).Methods(http.MethodGet)
     r.HandleFunc("/dashboard", handlers.GetDashboardHandler(sessionStore)).Methods(http.MethodGet)
}


func StartServer() error {
     appConfig, err := config.GetConfig()

     if err != nil {
          return err
     }

     conString := connectionString(appConfig)

     err = db.Run(conString)

     if err != nil {
          return err
     }

     sessionStore := providers.NewCookieStore(handlers.CookieName)
     r := newRouter()
     r.setupRoutes(conString, sessionStore)
     log.Printf("Server starting at port ...%s", appConfig.Port)
     if err := http.ListenAndServe(":"+appConfig.Port, r); err!= nil {
          return err
     }

     return nil
}

func connectionString(appConfig config.AppConfig) string {
   return   fmt.Sprintf( "host=%s port=%s user=%s dbname=%s password=%s sslmode=%s", appConfig.DBHost, appConfig.DBPort, appConfig.DBUser, appConfig.DBName, appConfig.DBPassword, appConfig.SSLMode )
}