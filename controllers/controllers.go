package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AntonioGSC/Api-Rest-Go/database"
	"github.com/AntonioGSC/Api-Rest-Go/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade

	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}

func CriarPersonalidade(w http.ResponseWriter, r *http.Request) {
	var p models.Personalidade
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Create(&p)
	json.NewEncoder(w).Encode(p)
}

func DeletaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade

	database.DB.Delete(&p, id)
	json.NewEncoder(w).Encode(p)
}

func EditaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalidade

	database.DB.First(&p, id)
	json.NewDecoder(r.Body).Decode(&p)
	database.DB.Save(&p)
	json.NewEncoder(w).Encode(p)
}
