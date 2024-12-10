package controllers

import (
    "armstrong-webapp/database"
    "armstrong-webapp/models"
    "encoding/json"
    "net/http"
)

func RegisterUser(w http.ResponseWriter, r *http.Request) {
    // Ensure the method is POST
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    // Parse input JSON request
    var input struct {
        Email string `json:"email"`
    }

    if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    // Create a new User object
    user := models.User{
        Email: input.Email,
    }

    // Insert the user into the database
    result := database.DB.Create(&user)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    // Prepare the response with the user_id
    response := map[string]interface{}{
        "message": "User registered successfully",
        "user_id": user.ID, // Assuming ID is the primary key
    }

    // Set header and encode the response as JSON
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated) // 201 Created
    json.NewEncoder(w).Encode(response)
}
