package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"cambria-bank/handlers"
	"cambria-bank/middleware"
)

func SetupRoutes() *mux.Router {

	router := mux.NewRouter()

	// Serve CSS, JavaScript and other static files
	router.PathPrefix("/static/").Handler(
		http.StripPrefix(
			"/static/",
			http.FileServer(http.Dir("static")),
		),
	)

	// Login page
	router.HandleFunc("/login-page", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/login.html")
	}).Methods("GET")

	// Registration page
	router.HandleFunc("/register-page", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/register.html")
	}).Methods("GET")

	// Dashboard page
	router.HandleFunc("/dashboard", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "templates/dashboard.html")
	}).Methods("GET")

	// API routes
	router.HandleFunc("/register", handlers.Register).Methods("POST")
	router.HandleFunc("/login", handlers.Login).Methods("POST")

	// Protected account API
	router.Handle(
		"/account",
		middleware.Auth(http.HandlerFunc(handlers.GetAccount)),
	).Methods("GET")
	// Protected transaction history API
	router.Handle(
		"/transactions",
		middleware.Auth(http.HandlerFunc(handlers.GetTransactions)),
	).Methods("GET")

	// Home page
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/login-page", http.StatusFound)
	}).Methods("GET")

	return router
}
