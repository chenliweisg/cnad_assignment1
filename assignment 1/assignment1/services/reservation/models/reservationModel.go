package models

import (
	reservation_db "assignment1/services/reservation/utils" // imports the utils package from the reservation directory in your project
	vehicles_db "assignment1/services/vehicles/utils"
	// "log"
	//"time"
)

// Reservation represents the Reservations table
type Reservation struct {
	ReservationID int    `json:"reservationID"` // Maps to ReservationID in the database
	UserID        int    `json:"userID"`        // Maps to UserID in the database
	VehicleID     int    `json:"vehicleID"`     // Maps to VehicleID in the database
	StartTime     string `json:"startTime"`     // Maps to StartTime in the database
	EndTime       string `json:"endTime"`       // Maps to EndTime in the database
	Status        string `json:"status"`        // Maps to Status in the database
}

func CreateReservation(reservation Reservation) error {
	// Step 1: Start transactions for both databases
	res_db := reservation_db.GetDB()
	veh_db := vehicles_db.GetDB()

	res_tx, err := res_db.Begin()
	if err != nil {
		return err
	}

	veh_tx, err := veh_db.Begin()
	if err != nil {
		res_tx.Rollback()
		return err
	}

	// Step 2: Insert reservation into reservation_db
	_, err = res_tx.Exec(`
        INSERT INTO Reservations (UserID, VehicleID, StartTime, EndTime, Status)
        VALUES (?, ?, ?, ?, ?)`,
		reservation.UserID, reservation.VehicleID, reservation.StartTime, reservation.EndTime, reservation.Status,
	)
	if err != nil {
		res_tx.Rollback()
		veh_tx.Rollback()
		return err
	}

	// Step 3: Update vehicle status in vehicles_db
	vstatus := "" // Initialize vstatus
	if reservation.Status == "Booked" {
		vstatus = "Unavailable"
	}

	_, err = veh_tx.Exec(`
        UPDATE Vehicles SET Status = ? WHERE VehicleID = ?`,
		vstatus, reservation.VehicleID,
	)
	if err != nil {
		res_tx.Rollback()
		veh_tx.Rollback()
		return err
	}

	// Step 4: Commit both transactions
	if err := res_tx.Commit(); err != nil {
		veh_tx.Rollback()
		return err
	}
	if err := veh_tx.Commit(); err != nil {
		return err
	}

	return nil
}

// FetchReservationsByUser retrieves all reservations for a specific user
func FetchReservationsByUser(userID int) ([]Reservation, error) {
	db := reservation_db.GetDB()

	rows, err := db.Query(`
		SELECT ReservationID, UserID, VehicleID, StartTime, EndTime, Status
		FROM Reservations
		WHERE UserID = ? AND Status = "Booked"`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []Reservation
	for rows.Next() {
		var reservation Reservation
		if err := rows.Scan(&reservation.ReservationID, &reservation.UserID, &reservation.VehicleID, &reservation.StartTime, &reservation.EndTime, &reservation.Status); err != nil {
			return nil, err
		}
		reservations = append(reservations, reservation)
	}
	return reservations, nil
}

// UpdateReservation modifies a reservation's start and end time
func UpdateReservation(reservationID int, updatedData Reservation) error {
	db := reservation_db.GetDB()

	_, err := db.Exec(`
		UPDATE Reservations
		SET StartTime = ?, EndTime = ?
		WHERE ReservationID = ?`,
		updatedData.StartTime, updatedData.EndTime, reservationID)
	return err
}

// GetVehicleIDByReservation fetches the VehicleID for a given reservation
func GetVehicleIDByReservation(reservationID int) (int, error) {
	db := reservation_db.GetDB()

	var vehicleID int
	err := db.QueryRow(`
        SELECT VehicleID FROM Reservations
        WHERE ReservationID = ?`, reservationID).Scan(&vehicleID)
	if err != nil {
		return 0, err
	}

	return vehicleID, nil
}

// UpdateVehicleStatus updates the status of a vehicle in the Vehicles table
func UpdateVehicleStatus(vehicleID int, status string) error {
	db := vehicles_db.GetDB()

	_, err := db.Exec(`
        UPDATE Vehicles
        SET Status = ?
        WHERE VehicleID = ?`, status, vehicleID)
	return err
}

// DeleteReservation deletes a reservation by ID
func DeleteReservation(reservationID int) error {
	db := reservation_db.GetDB()

	_, err := db.Exec(`
		UPDATE Reservations
		SET Status = "Canceled"
		WHERE ReservationID = ?`, reservationID)
	return err
}

// Return a vehicle
func ReturnReservation(reservationID int, vehicleID int) error {
	reservationDB := reservation_db.GetDB()
	vehicleDB := vehicles_db.GetDB()

	// Update reservation status to 'Complete'
	_, err := reservationDB.Exec(`UPDATE Reservations SET Status = 'Complete' WHERE ReservationID = ?`, reservationID)
	if err != nil {
		return err
	}

	// Update vehicle status to 'Available'
	_, err = vehicleDB.Exec(`UPDATE Vehicles SET Status = 'Available' WHERE VehicleID = ?`, vehicleID)
	if err != nil {
		return err
	}

	return nil
}

// FetchReservationHistory retrieves reservation history for a specific user
func FetchReservationHistory(userID string) ([]Reservation, error) {
	db := reservation_db.GetDB()

	rows, err := db.Query(`
        SELECT ReservationID, VehicleID, StartTime, EndTime, Status
        FROM Reservations
        WHERE UserID = ? AND Status != 'Booked'`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reservations []Reservation
	for rows.Next() {
		var reservation Reservation
		if err := rows.Scan(&reservation.ReservationID, &reservation.VehicleID, &reservation.StartTime, &reservation.EndTime, &reservation.Status); err != nil {
			return nil, err
		}
		reservations = append(reservations, reservation)
	}

	return reservations, nil
}
