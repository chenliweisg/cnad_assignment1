package controllers

import (
	"assignment1/services/reservation/models" // Import reservation model into here
	"encoding/json"
	"net/http"
	"strconv"

	//"log"

	"github.com/gorilla/mux"
)

// Book a vehicle reservation
func CreateReservation(w http.ResponseWriter, r *http.Request) {
	var reservation models.Reservation
	err := json.NewDecoder(r.Body).Decode(&reservation)
	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = models.CreateReservation(reservation)
	if err != nil {
		http.Error(w, "Failed to create reservation: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Reservation created successfully"))
}

// GetReservationsByUser retrieves all reservations for a specific user
func GetReservationsByUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, err := strconv.Atoi(vars["userID"])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	reservations, err := models.FetchReservationsByUser(userID)
	if err != nil {
		http.Error(w, "Failed to fetch reservations: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reservations)
}

// ModifyReservation modifies an existing reservation
func ModifyReservation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reservationID, err := strconv.Atoi(vars["reservationID"])
	if err != nil {
		http.Error(w, "Invalid reservation ID", http.StatusBadRequest)
		return
	}

	var updatedData models.Reservation
	err = json.NewDecoder(r.Body).Decode(&updatedData)
	if err != nil {
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	err = models.UpdateReservation(reservationID, updatedData)
	if err != nil {
		http.Error(w, "Failed to modify reservation: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Reservation modified successfully"))
}

// CancelReservation cancels an existing reservation and updates the vehicle status
func CancelReservation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reservationID, err := strconv.Atoi(vars["reservationID"])
	if err != nil {
		http.Error(w, "Invalid reservation ID", http.StatusBadRequest)
		return
	}

	// Fetch the VehicleID for the reservation being canceled
	vehicleID, err := models.GetVehicleIDByReservation(reservationID)
	if err != nil {
		http.Error(w, "Failed to fetch vehicle ID: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Cancel the reservation
	err = models.DeleteReservation(reservationID)
	if err != nil {
		http.Error(w, "Failed to cancel reservation: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Update the vehicle status to "Available"
	err = models.UpdateVehicleStatus(vehicleID, "Available")
	if err != nil {
		http.Error(w, "Failed to update vehicle status: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Reservation canceled and vehicle status updated successfully"))
}

// Return Reservation
func ReturnReservation(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reservationID, err := strconv.Atoi(vars["reservationID"])
	if err != nil {
		http.Error(w, "Invalid reservation ID", http.StatusBadRequest)
		return
	}

	var payload struct { //correctly receive and handle the vehicleID sent in the AJAX request.
		VehicleID int `json:"vehicleID"`
	}

	//log.Printf("Received VehicleID: %d", payload.VehicleID)

	// Parse the JSON body
	err = json.NewDecoder(r.Body).Decode(&payload)
	if err != nil || payload.VehicleID == 0 {
		http.Error(w, "Invalid vehicle ID", http.StatusBadRequest)
		return
	}

	vehicleID := payload.VehicleID

	// Call the model to update the database
	err = models.ReturnReservation(reservationID, vehicleID)
	if err != nil {
		http.Error(w, "Failed to return reservation: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Reservation and vehicle successfully returned"))
}

// GetReservationHistory retrieves reservation history for a specific user
func GetReservationHistory(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]

	reservations, err := models.FetchReservationHistory(userID)
	if err != nil {
		http.Error(w, "Failed to fetch reservation history: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(reservations)
}
