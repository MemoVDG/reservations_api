package resolvers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/MemoVDG28/reservations_api/configs"
	"github.com/MemoVDG28/reservations_api/helpers"
	"github.com/MemoVDG28/reservations_api/models"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

var DB, _ = configs.ConnectDB()
var reservationCollection *mongo.Collection = configs.GetCollection(DB, "reservations")

func CreateReservationResolver() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		reservation := models.Reservation{}
		// TODO: Implement marshal for dates
		// TODO: Validations
		//body, errr := ioutil.ReadAll(r.Body)
		//if errr != nil {
		//	panic(errr)
		//}
		//errx := json.Unmarshal(body, &reservation)
		//if errx != nil {
		//	panic(errx)
		//}
		err := json.NewDecoder(r.Body).Decode(&reservation)
		newReservation := models.Reservation{
			Id:        primitive.NewObjectID(),
			User:      reservation.User,
			Property:  reservation.Property,
			StartDate: reservation.StartDate,
			EndDate:   reservation.EndDate,
		}

		result, err := reservationCollection.InsertOne(ctx, newReservation)
		if err != nil {
			var response = map[string]string{
				"error": err.Error(),
			}
			helpers.SendResponse(rw, r, response, http.StatusInternalServerError)
			return
		}
		helpers.SendResponse(rw, r, result, http.StatusInternalServerError)
	}
}

func GetReservationById() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		params := mux.Vars(request)
		reservationId := params["reservationId"]
		defer cancel()

		reservation := models.Reservation{}
		objId, _ := primitive.ObjectIDFromHex(reservationId)

		err := reservationCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&reservation)
		if err != nil {
			var response = map[string]string{
				"error": err.Error(),
			}
			helpers.SendResponse(writer, request, response, http.StatusInternalServerError)
			return
		}

		helpers.SendResponse(writer, request, reservation, http.StatusOK)
		fmt.Println("Success")

	}
}
