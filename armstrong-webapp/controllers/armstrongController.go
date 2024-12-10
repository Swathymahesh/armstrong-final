package controllers

import (
	"armstrong-webapp/database"
	"armstrong-webapp/models"
	"encoding/json"
	       // For formatted output (optional, if needed)
	"io"         // For reading the raw body
	"log"        // For logging debug information
	"net/http"
	"strconv"
	"strings"
)


// VerifyArmstrongNumber checks if a number is an Armstrong number and saves it for a user
// Debug log to inspect the incoming request body
func VerifyArmstrongNumber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var input struct {
		UserID uint `json:"user_id"`
		Number int  `json:"number"`
	}

	// Debug: Log the raw body
	body := new(strings.Builder)
	_, err := io.Copy(body, r.Body)
	if err != nil {
		http.Error(w, "Error reading body", http.StatusBadRequest)
		return
	}
	log.Println("Raw Body:", body.String())

	// Decode the JSON
	if err := json.NewDecoder(strings.NewReader(body.String())).Decode(&input); err != nil {
		log.Println("JSON Decoding Error:", err)
		http.Error(w, "Invalid input: "+err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("Decoded Input: UserID=%d, Number=%d\n", input.UserID, input.Number)

	// Validate inputs
	if input.UserID == 0 || input.Number <= 0 {
		http.Error(w, "Invalid input: user_id and number are required", http.StatusBadRequest)
		return
	}

	// Armstrong number logic
	num, temp, sum := input.Number, input.Number, 0
	for temp > 0 {
		digit := temp % 10
		sum += digit * digit * digit
		temp /= 10
	}

	if num != sum {
		http.Error(w, "Not an Armstrong number", http.StatusBadRequest)
		return
	}

	// Save to database
	armstrong := models.ArmstrongNumber{
		UserID: input.UserID,
		Number: num,
	}
	result := database.DB.Create(&armstrong)
	if result.Error != nil {
		http.Error(w, "Database error: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
response := map[string]string{"message": fmt.Sprintf("%d is an Armstrong number and has been saved", num)}
json.NewEncoder(w).Encode(response)

}

// GetUserArmstrongNumbers retrieves Armstrong numbers for a specific user
func GetUserArmstrongNumbers(w http.ResponseWriter, r *http.Request) {
	// Ensure it's a GET request
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Extract user ID from the URL
	pathParts := strings.Split(strings.TrimPrefix(r.URL.Path, "/user/"), "/numbers")
	if len(pathParts) == 0 || pathParts[0] == "" {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(pathParts[0])
	if err != nil || userID <= 0 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Fetch Armstrong numbers from the database
	var numbers []models.ArmstrongNumber
	result := database.DB.Where("user_id = ?", userID).Find(&numbers)
	if result.Error != nil {
		http.Error(w, "Database error: "+result.Error.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the results
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(numbers)
}

// GetAllUsersAndNumbers retrieves all users and their Armstrong numbers
func GetAllUsersAndNumbers(w http.ResponseWriter, r *http.Request) {
	// Ensure it's a GET request
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Fetch data with a JOIN query
	var results []struct {
		Email  string `json:"email"`
		Number int    `json:"number"`
	}
	database.DB.Raw(`
		SELECT users.email, armstrong_numbers.number
		FROM users
		JOIN armstrong_numbers ON users.id = armstrong_numbers.user_id
	`).Scan(&results)

	// Respond with the results
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
