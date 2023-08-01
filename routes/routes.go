package routes

import (
	"challenge-milhas/controllers"
	"challenge-milhas/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequests() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/depoimentos", controllers.GetAllDepoimentos).Methods("Get")
	r.HandleFunc("/api/depoimentos/{id}", controllers.GetDepoimento).Methods("Get")
	r.HandleFunc("/api/depoimentos", controllers.CreateDepoimento).Methods("Post")
	r.HandleFunc("/api/depoimentos/{id}", controllers.DeleteDepoimentoo).Methods("Delete")
	r.HandleFunc("/api/depoimentos/{id}", controllers.UpdateDepoimento).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", r))
}
