package models

import (
	"assignment1/services/user/utils"
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Phone        string `json:"phone"`
	Password     string `json:"password"` // Change from PasswordHash to Password
	MembershipID int    `json:"membershipID"`
}

// FetchAllUsers retrieves all users from the database
func FetchAllUsers() ([]User, error) {
	db := utils.GetDB() // Ensure connection is valid

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Close rows after processing

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Password, &user.MembershipID); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// HashPassword hashes the password using bcrypt
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return "", err
	}
	return string(hashedPassword), nil
}

// CreateUser inserts a new user into the database with a hashed password
func CreateUser(user User) error {
	db := utils.GetDB()
	log.Println("memershiperID:", user.MembershipID)

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Insert user into the database
	_, err = db.Exec(`
        INSERT INTO Users (Name, Email, Phone, PasswordHash, MembershipID)
        VALUES (?, ?, ?, ?, ?)`,
		user.Name, user.Email, user.Phone, string(hashedPassword), user.MembershipID)
	return err
}

// ValidateUser checks if the email and password are correct for login
func ValidateUser(email, password string) (int, int, bool, error) {
	db := utils.GetDB() // Ensure connection is valid

	var userID int
	var membershipID int
	var storedPasswordHash string

	// Fetch user details
	err := db.QueryRow("SELECT userID, membershipID, passwordHash FROM users WHERE email = ?", email).
		Scan(&userID, &membershipID, &storedPasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("No user found with email:", email)
			return 0, 0, false, nil // User not found
		}
		log.Println("Error retrieving user:", err)
		return 0, 0, false, err // Database query error
	}

	// Compare the provided password with the stored hash
	err = bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(password))
	if err != nil {
		log.Println("Password mismatch for email:", email)
		return 0, 0, false, nil // Passwords do not match
	}

	return userID, membershipID, true, nil // Credentials are valid
}

// GetUserByID fetches a user by their ID
func GetUserByID(userID string) (*User, error) {
	db := utils.GetDB()

	var user User
	err := db.QueryRow("SELECT userid, name, email, phone, passwordHash, membershipID FROM users WHERE userid = ?", userID).Scan(&user.ID, &user.Name, &user.Email, &user.Phone, &user.Password, &user.MembershipID)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// UpdateUser updates the user's information in the database
func UpdateUser(user User) error {
	db := utils.GetDB()

	// Initialize variables for hashed password
	var hashedPassword string
	var err error

	// If the password field is not empty, hash the new password
	if user.Password != "" {
		hashedPasswordBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("Error hashing password: %v\n", err)
			return err
		}
		hashedPassword = string(hashedPasswordBytes)
	}

	// Prepare the SQL query based on whether a password update is needed
	var query string
	var args []interface{}

	if user.Password != "" {
		// Update including the new hashed password
		query = `
			UPDATE Users
			SET Name = ?, Email = ?, Phone = ?, PasswordHash = ?, MembershipID = ?
			WHERE UserID = ?`
		args = []interface{}{user.Name, user.Email, user.Phone, hashedPassword, user.MembershipID, user.ID}
	} else {
		// Update without changing the password
		query = `
			UPDATE Users
			SET Name = ?, Email = ?, Phone = ?, MembershipID = ?
			WHERE UserID = ?`
		args = []interface{}{user.Name, user.Email, user.Phone, user.MembershipID, user.ID}
	}

	// Execute the query
	_, err = db.Exec(query, args...)
	if err != nil {
		log.Printf("Error updating user: %v\n", err)
		return err
	}

	return nil
}

func FetchDiscountRate(membershipID int) (float64, error) {
	db := utils.GetDB()

	var discountRate float64
	err := db.QueryRow(`
        SELECT DiscountRate
        FROM MembershipTiers
        WHERE MembershipID = ?`, membershipID).Scan(&discountRate)
	if err != nil {
		return 0, err
	}

	return discountRate, nil
}
