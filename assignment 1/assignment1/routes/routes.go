package routes

import (
	billingControllers "assignment1/services/billing/controllers"
	reservationControllers "assignment1/services/reservation/controllers"
	userControllers "assignment1/services/user/controllers"
	vehicleControllers "assignment1/services/vehicles/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

// SetupRoutes defines all API routes
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// User Routes
	router.HandleFunc("/api/v1/users", userControllers.GetAllUsers).Methods("GET")

	//Create a new user
	router.HandleFunc("/api/v1/signup", userControllers.Signup).Methods("POST")

	//Handles user login requests
	router.HandleFunc("/api/v1/login", userControllers.Login).Methods("POST")

	// Serve index.html at "/login"
	router.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/login.html")
	}).Methods("GET")

	//******TO USER CONTROLLER******
	// User routes
	router.HandleFunc("/api/v1/profile/{userid}", userControllers.GetUserProfile).Methods("GET")

	//URL to updates user information
	router.HandleFunc("/api/v1/profile/{userid}", userControllers.UpdateUserProfile).Methods("PUT")

	// URL to get Discount Rate for calculate cost in reservation page
	router.HandleFunc("/api/v1/membership/{membershipID}/discount", userControllers.GetMembershipDiscount).Methods("GET")

	// ******TO VEHICLE CONTROLLER******
	// URL to get all Vehicle that is available
	router.HandleFunc("/api/v1/vehicles/available", vehicleControllers.GetAvailableVehicles).Methods("GET")

	//******TO RESERVATION CONTOLLER******
	// URL to Book Reservation for a vehicle
	router.HandleFunc("/api/v1/reservations", reservationControllers.CreateReservation).Methods("POST")

	// Get all reservations for a specific user with 'book' status
	router.HandleFunc("/api/v1/reservations/user/{userID}", reservationControllers.GetReservationsByUser).Methods("GET")

	// Return a vechicle for reservation
	router.HandleFunc("/api/v1/reservations/{reservationID}/return", reservationControllers.ReturnReservation).Methods("PUT")

	// Modify a reservation
	router.HandleFunc("/api/v1/reservations/{reservationID}", reservationControllers.ModifyReservation).Methods("PUT")

	// Cancel a reservation
	router.HandleFunc("/api/v1/reservations/{reservationID}", reservationControllers.CancelReservation).Methods("DELETE")

	// Get all reservation hisotry for a specific user
	router.HandleFunc("/api/v1/reservations/history/{userID}", reservationControllers.GetReservationHistory).Methods("GET")


	//******TO Billing CONTOLLER******
	// Create a transaction after payment
	router.HandleFunc("/api/v1/transactions", billingControllers.CreateTransaction).Methods("POST")
	// Get invoice for a user
	router.HandleFunc("/api/v1/invoices/{userID}", billingControllers.GetUserInvoices).Methods("GET")

	return router
}
