

USE vehicles_db;
select * from Vehicles;

USE reservation_db;
select * from Reservations;

USE billing_db;
select * from Transactions;
select * from Invoices;

use user_db;
select * from users;
SELECT passwordHash FROM users WHERE email = 'john.doe@example.com';
UPDATE users
        SET name = 'John Doe', email = 'john.doe@example.com', phone = '1234567890', passwordHash = 'password123', membershipID = '1'
        WHERE userid = 1;
SELECT userid, name, email, phone, passwordHash, membershipID FROM users WHERE userid = "5";


INSERT INTO Reservations (UserID, VehicleID, StartTime, EndTime, Status)
VALUES
    (1, 3, '2024-12-06 08:00:00', '2024-12-06T17:05', 'Active')
    
