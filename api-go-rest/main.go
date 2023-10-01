package main

import ( 
	"fmt"
	"api-go-rest/routes"
	"api-go-rest/models"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Leo", Historia: "test"},
		{Id: 2, Nome: "Leo2", Historia: "test2"},
	}


	fmt.Println("Iniciando o servidor Rest com GO")
	routes.HandleRequest()
}