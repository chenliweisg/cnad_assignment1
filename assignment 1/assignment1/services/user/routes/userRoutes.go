package routes

import (
	"assignment1/services/user/controllers"

	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router) {
	router.HandleFunc("/api/v1/users", controllers.GetAllUsers).Methods("GET")
}
