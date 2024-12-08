package models

import ( //"assignment1/services/vehicles/utils"
	// reservation_db "assignment1/services/reservation/utils"
	vehicles_db "assignment1/services/vehicles/utils"
)

type Vehicle struct {
	VehicleID   int     `json:"vehicleID"`   // Maps to VehicleID in the database
	Model       string  `json:"model"`       // Maps to Model in the database
	Location    string  `json:"location"`    // Maps to Location in the database
	ChargeLevel float64 `json:"chargeLevel"` // Maps to ChargeLevel in the database
	Status      string  `json:"status"`      // Maps to Status in the database
}

// func getLocalTime() time.Time {
// 	// Replace "Asia/Singapore" with your timezone (e.g., "America/New_York", "UTC")
// 	location, err := time.LoadLocation("Asia/Singapore")
// 	if err != nil {
// 		log.Fatalf("Failed to load timezone: %v", err)
// 	}
// 	return time.Now().In(location)
// }

// func UpdateExpiredReservationsAndVehicles() error {

// 	// currentTimeUTC := currentTime.UTC()
// 	currentTime := getLocalTime().Format("2006-01-02 15:04:05")

// 	// Log the current local time for debugging
// 	log.Printf("Current Local Time: %v", currentTime)

// 	// Step 1: Update expired reservations in reservation_db
// 	res_db := reservation_db.GetDB()
// 	veh_db := vehicles_db.GetDB()

// 	// Find expired reservations with 'Active' status
// 	rows, err := res_db.Query(`
//         SELECT VehicleID FROM reservations
//         WHERE EndTime <= ? AND Status = 'Booked'`,
// 		currentTime,
// 	)
// 	if err != nil {
// 		return err
// 	}
// 	defer rows.Close()

// 	// Collect VehicleIDs for which reservations are updated
// 	var vehicleIDs []int
// 	for rows.Next() {
// 		var vehicleID int
// 		if err := rows.Scan(&vehicleID); err != nil {
// 			return err
// 		}
// 		vehicleIDs = append(vehicleIDs, vehicleID)
// 	}
// 	log.Printf("Number of Expired Vehicle IDs: %d", len(vehicleIDs))

// 	// Step 2: Update reservation statuses to 'Complete'
// 	_, err = res_db.Exec(`
// 		UPDATE Reservations
// 		SET Status = 'Completed'
// 		WHERE ReservationID IN (
// 			SELECT ReservationID
// 			FROM (
// 				SELECT ReservationID
// 				FROM Reservations
// 				WHERE EndTime <= ? AND Status = 'Booked'
// 			) AS temp
// 		)`,
// 		currentTime,
// 	)
// 	if err != nil {
// 		return err
// 	}

// 	// Step 3: Update vehicle statuses to 'Available' in vehicles_db
// 	for _, vehicleID := range vehicleIDs {
// 		_, err = veh_db.Exec(`
//             UPDATE Vehicles SET Status = 'Available'
//             WHERE VehicleID = ?`,
// 			vehicleID,
// 		)
// 		if err != nil {
// 			log.Printf("Failed to update vehicle %d status: %v", vehicleID, err)
// 		} else {
// 			log.Printf("Vehicle %d marked as Available", vehicleID)
// 		}
// 	}

// 	return nil
// }

func GetVehiclesByStatus(status string) ([]Vehicle, error) {
	db := vehicles_db.GetDB()

	rows, err := db.Query("SELECT VehicleID, Model, Location, ChargeLevel, Status FROM Vehicles WHERE Status = ?", status)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vehicles []Vehicle
	for rows.Next() {
		var vehicle Vehicle
		if err := rows.Scan(&vehicle.VehicleID, &vehicle.Model, &vehicle.Location, &vehicle.ChargeLevel, &vehicle.Status); err != nil {
			return nil, err
		}
		vehicles = append(vehicles, vehicle)
	}
	return vehicles, nil
}

func GetVehicleModel(vehicleID int) (string, error) {
	db := vehicles_db.GetDB()

	var vehicleModel string
	err := db.QueryRow(`
        SELECT Model
        FROM Vehicles
        WHERE VehicleID = ?`, vehicleID).Scan(&vehicleModel)
	if err != nil {
		return "", err
	}
	return vehicleModel, nil
}
