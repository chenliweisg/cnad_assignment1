package controllers

import (
	"assignment1/services/user/models"
	"encoding/json"
	"net/http"
	"strconv"

	"log"

	"github.com/gorilla/mux"
)

// GetAllUsers handles GET /api/v1/users
func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := models.FetchAllUsers()
	if err != nil {
		http.Error(w, "Error fetching users: "+err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Login Struct
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Signup handles user registration
func Signup(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	// Parse JSON body
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Create the user in the database
	err = models.CreateUser(newUser)
	if err != nil {
		http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

// Func to login
func Login(w http.ResponseWriter, r *http.Request) {
	var loginReq LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Validate user credentials and fetch userID and membershipID
	userID, membershipID, valid, err := models.ValidateUser(loginReq.Email, loginReq.Password)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}
	if !valid {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	// Return userID and membershipID as JSON
	response := map[string]interface{}{
		"userID":       userID,
		"membershipID": membershipID,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// GetUserProfile retrieves the logged-in user's profile
func GetUserProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) // Extract path variables
	userID := vars["userid"]

	// Fetch user from the database
	user, err := models.GetUserByID(userID)
	if err != nil {
		http.Error(w, "Failed to retrieve user profile: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Send user profile as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// updates the user profile
func UpdateUserProfile(w http.ResponseWriter, r *http.Request) {

	// Parse the JSON body
	var updatedUser models.User
	err := json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil {
		log.Printf("Error decoding JSON input: %v\n", err)
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Parsed user data: %+v\n", updatedUser)

	// Update the user in the database
	err = models.UpdateUser(updatedUser)
	if err != nil {
		log.Printf("Error updating user in database: %v\n", err)
		http.Error(w, "Failed to update user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User updated successfully"))
}

func GetMembershipDiscount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //	// Extract user ID from the URL
	membershipID, err := strconv.Atoi(vars["membershipID"])
	if err != nil {
		http.Error(w, "Invalid Membership ID", http.StatusBadRequest)
		return
	}

	discountRate, err := models.FetchDiscountRate(membershipID)
	if err != nil {
		http.Error(w, "Failed to fetch discount rate: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]float64{"discountRate": discountRate})
}
