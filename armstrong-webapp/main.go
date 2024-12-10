package main

import (
	"armstrong-webapp/database"
	"armstrong-webapp/routes"
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"math"
	"net/http"
)

func main() {
	database.ConnectDB()

	// Create a new router
	router := mux.NewRouter()

	// Add a default route
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Armstrong Web App is running!"))
	})

	// Register application routes
	routes.RegisterRoutes(router)

	// Apply CORS middleware
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}), // React app origin
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"}), // Allowed HTTP methods
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}), // Allowed headers
	)(router)

	// Start the server
	log.Println("Starting server on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", corsHandler))
}

// RegisterHandler handles user registration requests
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3001") // Match your frontend URL
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodPost {
		var data struct {
			Email string `json:"email"`
		}
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		response := map[string]string{"message": "User registered successfully"}
		json.NewEncoder(w).Encode(response)
		return
	}

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent) // Preflight request
		return
	}

	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
}

// VerifyHandler handles Armstrong number verification requests
func VerifyHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000") // Allow all origins
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodPost {
		var data struct {
			UserID string `json:"user_id"`
			Number int    `json:"number"`
		}
		if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Armstrong number logic
		isArmstrong := IsArmstrongNumber(data.Number)
		var message string
		if isArmstrong {
			message = "Number is an Armstrong number"
		} else {
			message = "Number is not an Armstrong number"
		}

		response := map[string]string{"message": message}
		json.NewEncoder(w).Encode(response)
		return
	}

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent) // Preflight request
		return
	}

	http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
}

// IsArmstrongNumber checks if a number is an Armstrong number
func IsArmstrongNumber(num int) bool {
	temp := num
	var digits []int

	// Extract digits
	for temp > 0 {
		digits = append(digits, temp%10)
		temp /= 10
	}

	// Calculate the Armstrong sum
	power := len(digits)
	sum := 0
	for _, digit := range digits {
		sum += int(math.Pow(float64(digit), float64(power)))
	}

	return sum == num
}
