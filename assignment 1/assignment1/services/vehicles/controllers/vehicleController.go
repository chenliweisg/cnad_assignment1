package controllers

import (
	"assignment1/services/vehicles/models"
	"encoding/json"
	"net/http"
	// "github.com/gorilla/mux"
)

// func GetAvailableVehicles(w http.ResponseWriter, r *http.Request) {
// 	vehicles, err := models.GetVehiclesByStatus("available")
// 	if err != nil {
// 		http.Error(w, "Failed to fetch vehicles: "+err.Error(), http.StatusInternalServerError)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(vehicles)
// 	log.Printf("Received request to update user with ID:")
// 	w.WriteHeader(http.StatusOK)
// }

func GetAvailableVehicles(w http.ResponseWriter, r *http.Request) {
	// // Step 1: Update expired reservations and vehicles
	// err := models.UpdateExpiredReservationsAndVehicles()
	// if err != nil {
	// 	http.Error(w, "Failed to update vehicle statuses: "+err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// Step 2: Retrieve available vehicles
	vehicles, err := models.GetVehiclesByStatus("Available")
	if err != nil {
		http.Error(w, "Failed to fetch available vehicles: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Step 3: Return the list of available vehicles as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vehicles)
}
