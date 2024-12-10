package routes

import (
    "armstrong-webapp/controllers"
    
	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {
    router.HandleFunc("/register", controllers.RegisterUser).Methods("POST")
    router.HandleFunc("/verify", controllers.VerifyArmstrongNumber).Methods("POST")
    router.HandleFunc("/user/{id}/numbers", controllers.GetUserArmstrongNumbers).Methods("GET")
    router.HandleFunc("/users", controllers.GetAllUsersAndNumbers).Methods("GET")
}
