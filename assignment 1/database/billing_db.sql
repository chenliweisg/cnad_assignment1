-- Create the billing table
CREATE database billing_db;
USE billing_db;

-- Create the Transactions table
CREATE TABLE Transactions (
    TransactionID INT AUTO_INCREMENT PRIMARY KEY,
    ReservationID INT NOT NULL,
    UserID INT NOT NULL,
    Amount DECIMAL(10, 2) NOT NULL,
    PaymentMethod VARCHAR(50) NOT NULL,
    Status VARCHAR(50) NOT NULL
);

-- Create the Invoices table
CREATE TABLE Invoices (
    InvoiceID INT AUTO_INCREMENT PRIMARY KEY,
    TransactionID INT NOT NULL,
    GeneratedTimestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    InvoiceDetails TEXT NOT NULL,
    FOREIGN KEY (TransactionID) REFERENCES Transactions(TransactionID)
);

-- Insert dummy data into Transactions
INSERT INTO Transactions (ReservationID, UserID, Amount, PaymentMethod, Status)
VALUES
    (1, 1, 50.00, 'Credit Card', 'Completed'),
    (2, 2, 30.00, 'PayPal', 'Completed'),
    (3, 3, 70.00, 'Debit Card', 'Pending'),
    (4, 4, 20.00, 'Credit Card', 'Refunded'),
    (5, 5, 100.00, 'Bank Transfer', 'Completed');

-- Insert dummy data into Invoices
INSERT INTO Invoices (TransactionID, InvoiceDetails)
VALUES
    (1, 'Reservation: Tesla Model X. Amount Paid: $50.00. Payment Method: Credit Card.'),
    (2, 'Reservation: Toyota Prius. Amount Paid: $30.00. Payment Method: PayPal.'),
    (3, 'Reservation: BMW i3. Amount Pending: $70.00. Payment Method: Debit Card.'),
    (4, 'Reservation: Chevy Bolt. Amount Refunded: $20.00. Payment Method: Credit Card.'),
    (5, 'Reservation: Nissan Leaf. Amount Paid: $100.00. Payment Method: Bank Transfer.');
    

