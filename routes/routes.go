package routes

import (
	"log"
	"net/http"

	"github.com/AntonioGSC/Api-Rest-Go/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
