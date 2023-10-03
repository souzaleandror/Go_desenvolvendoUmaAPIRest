package controllers

import ( 
	"net/http"
	"fmt"
	"encoding/json"
	"api-go-rest/models"
	
	"github.com/gorilla/mux"
	"api-go-rest/database"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade

	database.DB.Find(&p)

	//json.NewEncoder(w).Encode(models.Personalidades)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	// for _, personalidade := range models.Personalidades {
	// 	if strconv.Itoa(personalidade.Id) == id {
	// 		json.NewEncoder(w).Encode(personalidade)
	// 	}
	// }

	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)

}

func criaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var novaPersonalidade models.Personalidade
	json.NewDecoder(r.body).Decode(&novaPersonalidade)
	database.DB.Create(&novaPersonalidade)
	json.NewDecoder(w).Econde(novaPersonalidade)
}

func deletaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.Create(&personalidade, id)
	json.NewDecoder(w).Econde(personalidade)
}

func editarUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewDecoder(r.body).Decode(&personalidade)
	database.DB.Save(&personalidade)
	json.NewDecoder(w).Econde(personalidade)
}