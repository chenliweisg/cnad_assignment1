package models

import (
	"assignment1/services/billing/utils"
	"fmt"
)

type Invoice struct {
	InvoiceID      int    `json:"invoiceID"`
	TransactionID  int    `json:"transactionID"`
	InvoiceDetails string `json:"invoiceDetails"`
}

// CreateTransaction inserts a new transaction into the database and returns the transaction ID.
func CreateTransaction(reservationID, userID int, amount float64, paymentMethod, status string) (int, error) {
	db := utils.GetDB()

	result, err := db.Exec(`
        INSERT INTO Transactions (ReservationID, UserID, Amount, PaymentMethod, Status)
        VALUES (?, ?, ?, ?, ?)`,
		reservationID, userID, amount, paymentMethod, status)
	if err != nil {
		return 0, err
	}

	transactionID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(transactionID), nil
}

func CreateInvoice(transactionID int, vehicleModel string, amount float64, paymentMethod string) error {
	db := utils.GetDB()

	// Generate InvoiceDetails
	invoiceDetails := fmt.Sprintf("Reservation: %s. Amount Paid: $%.2f. Payment Method: %s.", vehicleModel, amount, paymentMethod)

	// Insert into Invoices table
	_, err := db.Exec(`
        INSERT INTO Invoices (TransactionID, InvoiceDetails)
        VALUES (?, ?)`,
		transactionID, invoiceDetails)
	return err
}

// FetchInvoicesByUserID retrieves invoices for a specific user
func FetchInvoicesByUserID(userID string) ([]Invoice, error) {
	db := utils.GetDB()

	rows, err := db.Query(`
        SELECT i.InvoiceID, i.TransactionID, i.InvoiceDetails
        FROM Invoices i
        JOIN Transactions t ON i.TransactionID = t.TransactionID
        WHERE t.UserID = ?`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var invoices []Invoice
	for rows.Next() {
		var invoice Invoice
		if err := rows.Scan(&invoice.InvoiceID, &invoice.TransactionID, &invoice.InvoiceDetails); err != nil {
			return nil, err
		}
		invoices = append(invoices, invoice)
	}

	return invoices, nil
}
