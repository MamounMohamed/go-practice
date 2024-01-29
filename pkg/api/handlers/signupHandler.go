package handlers

import (
	"encoding/json"
	"net/http"
)

type SignupRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupResponse struct {
	Message string `json:"message"`
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Decode the JSON request body
	var request SignupRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&request); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Perform signup logic here (you may want to store the user in a database, etc.)

	// Respond with a JSON message
	response := SignupResponse{
		Message: "Signup successful",
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
func main() {
	// Define the signup endpoint
	http.HandleFunc("/signup", signupHandler)

	// Start the server on port 8080
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}
