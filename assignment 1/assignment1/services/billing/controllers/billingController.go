package controllers

import (
	billingModel "assignment1/services/billing/models"
	vehicleModels "assignment1/services/vehicles/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type TransactionRequest struct {
	ReservationID int     `json:"reservationID"`
	VehicleID     int     `json:"vehicleID"` // Include vehicleID directly
	UserID        int     `json:"userID"`
	Amount        float64 `json:"amount"`
	PaymentMethod string  `json:"paymentMethod"`
	Status        string  `json:"status"`
}

// CreateTransaction handles the creation of a transaction
func CreateTransaction(w http.ResponseWriter, r *http.Request) {
	var req TransactionRequest
	log.Println(1231231)

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	// Step 1: Fetch Vehicle Model from Vehicle Service using provided VehicleID
	vehicleModel, err := vehicleModels.GetVehicleModel(req.VehicleID)
	if err != nil {
		http.Error(w, "Failed to fetch vehicle model from vehicle service", http.StatusInternalServerError)
		return
	}

	log.Println(vehicleModel)
	log.Println("Step1 pass")

	// Step 2: Insert Transaction into Billing DB
	transactionID, err := billingModel.CreateTransaction(req.ReservationID, req.UserID, req.Amount, req.PaymentMethod, req.Status)
	if err != nil {
		http.Error(w, "Failed to create transaction", http.StatusInternalServerError)
		return
	}

	log.Println("Step2 pass")

	// Step 3: Insert Invoice into Billing DB
	err = billingModel.CreateInvoice(transactionID, vehicleModel, req.Amount, req.PaymentMethod)
	if err != nil {
		http.Error(w, "Failed to create invoice", http.StatusInternalServerError)
		return
	}

	log.Println("Step3 pass")

	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Transaction and Invoice created successfully")
}

// GetUserInvoices retrieves all invoices for a specific user
func GetUserInvoices(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["userID"]

	invoices, err := billingModel.FetchInvoicesByUserID(userID)
	if err != nil {
		http.Error(w, "Failed to fetch invoices: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(invoices)
}
