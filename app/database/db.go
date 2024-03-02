package database

import (
	"log"

	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
	"os"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")
	
	fmt.Printf("DB_HOST: %s\n", dbHost)
	fmt.Printf("DB_USER: %s\n", dbUser)
	fmt.Printf("DB_PASSWORD: %s\n", dbPassword)
	fmt.Printf("DB_NAME: %s\n", dbName)
	fmt.Printf("DB_PORT: %s\n", dbPort)
	
	stringDeConexao := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", dbHost, dbUser, dbPassword, dbName, dbPort)
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
