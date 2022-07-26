package routes

import (
	"github.com/MemoVDG28/reservations_api/resolvers"
	"github.com/gorilla/mux"
)

func ReservationRoutes(router *mux.Router) {
	router.HandleFunc("/reservation", resolvers.CreateReservationResolver()).Methods("POST")
	router.HandleFunc("/reservation/{reservationId}", resolvers.GetReservationById()).Methods("GET")
}
