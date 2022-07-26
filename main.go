package main

import (
	"github.com/MemoVDG28/reservations_api/configs"
	"github.com/MemoVDG28/reservations_api/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	//run database
	client, ctx := configs.ConnectDB()
	defer configs.DisconnectDB(client, ctx)
	//routes
	routes.ReservationRoutes(router)
	log.Fatal(http.ListenAndServe(":8000", router))
}
