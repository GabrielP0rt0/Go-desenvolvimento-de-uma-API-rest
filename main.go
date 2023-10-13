package main

import (
	"fmt"
	"go-rest-api/database"
	"go-rest-api/routes"
)

func main() {
	fmt.Println("Conectando ao banco de dados")
	database.ConectaComBancoDeDados()
	fmt.Println("Conectado com sucesso...")
	fmt.Println("Iniciando servidor go api rest")
	routes.HandleRequest()
}
