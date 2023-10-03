package database

import (
	"log"
	"gorm.io/gorm"
  "gorm.io/driver/postgres"
)

var DB *gorm.DB
var err error

func ConnectandoComBancoDeDados() {
	conexao := "user=postgres dbname=personalidades password=123456 host=localhost sslmode=disable"
	DB, err = gorm.Open(postgres.Open(conexao))

	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

}