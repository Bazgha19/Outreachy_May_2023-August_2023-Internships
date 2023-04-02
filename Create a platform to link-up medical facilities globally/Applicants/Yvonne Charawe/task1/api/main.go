package main

import (
	"fmt"
	"golang-crud-rest-api/database"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Load Configurations from config.json using Viper
	LoadAppConfig();
	// Initialize Database
	database.Connect(AppConfig.ConnectionString);
	//database.Migrate()

	// Initialize the router
	router := mux.NewRouter().StrictSlash(true);
	//Register Routes
	HospitalRoutes(router);
	//Allow CORS
	handler := cors.Default().Handler(router);
	// Start the server
	fmt.Printf("Starting Server on port %s", AppConfig.Port);
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", AppConfig.Port), handler));
}
