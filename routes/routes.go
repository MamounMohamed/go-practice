package routes

import (
	"github.com/gorilla/mux"
	"github.com/yourusername/yourprojectname/api/handlers"
)

func SetupRoutes(r *mux.Router) {
	// Handle signup endpoint
	r.HandleFunc("/signup", handlers.SignupHandler).Methods("POST")

	// Handle signin endpoint
	r.HandleFunc("/signin", handlers.SigninHandler).Methods("POST")

	// Add more routes as needed
}
