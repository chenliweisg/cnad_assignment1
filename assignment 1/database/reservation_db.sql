-- Create the Reservation table
CREATE database reservation_db;
USE reservation_db;

-- Create the Reservations table
CREATE TABLE Reservations (
	ReservationID INT AUTO_INCREMENT PRIMARY KEY,
    UserID INT NOT NULL,
    VehicleID INT NOT NULL,
    StartTime DATETIME NOT NULL,
    EndTime DATETIME NOT NULL,
    Status VARCHAR(50) NOT NULL
);


-- Create the BookingHistory table (optional)
-- CREATE TABLE BookingHistory (
--     HistoryID INT AUTO_INCREMENT PRIMARY KEY,
--     ReservationID INT NOT NULL,
-- 	StartTime DATETIME NOT NULL,
--     EndTime DATETIME NOT NULL,
--     Status VARCHAR(50) NOT NULL,
--     FOREIGN KEY (ReservationID) REFERENCES Reservations(ReservationID)
-- );

-- drop table BookingHistory

-- Insert dummy data into Reservations
INSERT INTO Reservations (UserID, VehicleID, StartTime, EndTime, Status)
VALUES
    (1, 1, '2024-12-01 08:00:00', '2024-12-01 10:00:00', 'Active'),
    (2, 3, '2024-12-02 14:00:00', '2024-12-02 16:00:00', 'Completed'),
    (3, 5, '2024-12-03 09:00:00', '2024-12-03 12:00:00', 'Active'),
    (4, 4, '2024-12-04 11:00:00', '2024-12-04 13:00:00', 'Canceled'),
    (5, 2, '2024-12-05 15:00:00', '2024-12-05 18:00:00', 'Booked');
    
    
USE reservation_db;
select * from Reservations;

SELECT VehicleID FROM Reservations
        WHERE EndTime < current_time() AND Status = 'Active';
        

		UPDATE Reservations
		SET Status = 'Complete'
		WHERE ReservationID IN (
			SELECT ReservationID
			FROM (
				SELECT ReservationID
				FROM Reservations
				WHERE EndTime < current_time() AND Status = 'Active'
			) AS temp
		)
        
UPDATE reservations
SET status = 'Complete'
WHERE reservationID = 57

USE reservation_db;
SELECT * FROM reservations
WHERE EndTime <= '2024-12-06 20:46:25.1456685 +0800 +08' AND Status = 'Active'

USE reservation_db;
DELETE FROM reservations WHERE ReservationID > 6




