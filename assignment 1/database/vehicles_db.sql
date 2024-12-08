-- Create the Vehicles table
CREATE database vehicles_db;
USE vehicles_db;

CREATE TABLE Vehicles (
    VehicleID INT AUTO_INCREMENT PRIMARY KEY,
    Model VARCHAR(100) NOT NULL,
    Location VARCHAR(100) NOT NULL,
    ChargeLevel DECIMAL(5, 2) NOT NULL,
    Status VARCHAR(50) NOT NULL
);

-- Create the VehicleStatusHistory table (optional)
-- CREATE TABLE VehicleStatusHistory (
--     StatusID INT AUTO_INCREMENT PRIMARY KEY,
--     VehicleID INT NOT NULL,
--     Timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
--     Status VARCHAR(50) NOT NULL,
--     FOREIGN KEY (VehicleID) REFERENCES Vehicles(VehicleID)
-- );

-- Insert dummy data into Vehicles
INSERT INTO Vehicles (Model, Location, ChargeLevel, Status)
VALUES
    ('Tesla Model X', 'Garage A', 80.50, 'Available'),
    ('Nissan Leaf', 'Garage B', 45.00, 'Maintenance'),
    ('Toyota Prius', 'Garage A', 90.75, 'Booked'),
    ('Chevy Bolt', 'Garage C', 60.00, 'Available'),
    ('BMW i3', 'Garage A', 30.00, 'Booked');

USE vehicles_db;
SELECT VehicleID, Model, Location, ChargeLevel, Status FROM Vehicles WHERE Status = 'available';
UPDATE vehicles set status = 'Available' where vehicleID > '0';
select * from vehicles



