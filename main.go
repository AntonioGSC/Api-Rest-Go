package main

import (
	"fmt"

	"github.com/AntonioGSC/Api-Rest-Go/models"
	"github.com/AntonioGSC/Api-Rest-Go/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
