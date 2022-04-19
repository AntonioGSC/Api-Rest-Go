package main

import (
	"fmt"

	"github.com/AntonioGSC/Api-Rest-Go/database"
	"github.com/AntonioGSC/Api-Rest-Go/routes"
)

func main() {
	database.ConectaComDB()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
