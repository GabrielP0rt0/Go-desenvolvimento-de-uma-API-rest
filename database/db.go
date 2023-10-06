package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "root"
	password = "root"
	dbName   = "root"
)

func ConectaComBancoDeDados() {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)
	DB, err = gorm.Open(postgres.Open(conn))
	if err != nil {
		log.Panic("Erro ao conectar ao banco de dados")
	}
}
