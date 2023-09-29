package main

import ( 
	"fmt"
	"api-go-rest/routes"
	"api-go-rest/models"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Nome: "Leo", Historia: "test"},
		{Nome: "Leo2", Historia: "test2"},
	}


	fmt.Println("Iniciando o servidor Rest com GO")
	routes.HandleRequest()
}