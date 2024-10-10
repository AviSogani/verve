package route

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"verve/controller"
)

func Init() {
	router := mux.NewRouter()
	router.HandleFunc("/api/verve/accept", controller.Accept).Methods("GET")
	router.Handle("/", router)

	// Listen on port 8085
	if err := http.ListenAndServe(":8085", router); err != nil {
		log.Println("Error starting server:", err)
	}
}
