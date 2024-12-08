package main

import (
	"assignment1/routes"
	billingUtils "assignment1/services/billing/utils"
	reservationUtils "assignment1/services/reservation/utils"
	userUtils "assignment1/services/user/utils"
	vehiclesUtils "assignment1/services/vehicles/utils"
	"log"
	"net/http"
)

func main() {

	//initialize and manage the lifecycle of the database connection - user
	userUtils.InitDB()              //Initializes the database connection when the application start
	defer userUtils.GetDB().Close() //Returns the global database connection to be used elsewhere in the application and Ensures the database connection is closed properly when the application exits.

	//initialize and manage the lifecycle of the database connection - vehicle
	vehiclesUtils.InitDB()
	defer vehiclesUtils.GetDB().Close()

	//initialize and manage the lifecycle of the database connection - reservation
	reservationUtils.InitDB()
	defer reservationUtils.GetDB().Close()

	//initialize and manage the lifecycle of the database connection - billing
	billingUtils.InitDB()
	defer billingUtils.GetDB().Close()

	// Initialize the router
	router := routes.SetupRoutes()

	// Serve static files through the router
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// Start the server
	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
