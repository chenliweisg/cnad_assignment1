#Create new database for user & create the table
CREATE database user_db;
USE user_db;

-- Create the Users table
CREATE TABLE Users (
    UserID INT AUTO_INCREMENT PRIMARY KEY,
    Name VARCHAR(100) NOT NULL,
    Email VARCHAR(100) UNIQUE NOT NULL,
    Phone VARCHAR(15) NOT NULL,
    PasswordHash VARCHAR(255) NOT NULL,
    MembershipID INT NOT NULL,
    FOREIGN KEY (MembershipID) REFERENCES MembershipTiers(MembershipID)
);

-- Create the MembershipTiers table
CREATE TABLE MembershipTiers (
    MembershipID INT AUTO_INCREMENT PRIMARY KEY,
    TierName VARCHAR(50) NOT NULL,
    DiscountRate DECIMAL(5, 2) NOT NULL,
    PriorityLevel INT NOT NULL, -- Lower number = higher priority
    BookingLimit INT NULL   -- Max number of active bookings NULL means no limit
);

-- drop table MembershipTiers,Users
-- Insert dummy data into MembershipTiers
INSERT INTO MembershipTiers (TierName, DiscountRate, PriorityLevel, BookingLimit)
VALUES
    ('Basic', 5.00, 3, 5),
    ('Premium', 10.00, 2, 10),
    ('VIP', 15.00, 1, NULL),
    ('Student', 7.50, 4, 3),
    ('Corporate', 20.00, 1, NULL);

-- Insert dummy data into Users
INSERT INTO Users (Name, Email, Phone, PasswordHash, MembershipID)
VALUES
    ('John Doe', 'john.doe@example.com', '1234567890', 'password123', 1),
    ('Jane Smith', 'jane.smith@example.com', '0987654321', 'securepass456', 2),
    ('Alice Johnson', 'alice.johnson@example.com', '1112223333', 'mypassword789', 3),
    ('Bob Brown', 'bob.brown@example.com', '4445556666', MD5('bobspassword'), 4),
    ('Charlie White', 'charlie.white@example.com', '7778889999', MD5('charliepass'), 5);


USE user_db;
select * from users;
select * from membershiptiers;

UPDATE users set membershipID = 5 where userID =18

       SELECT DiscountRate
        FROM MembershipTiers
        WHERE MembershipID = 5