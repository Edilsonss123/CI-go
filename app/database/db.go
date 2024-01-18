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
	bdHost := os.Getenv("BD_HOST")
	stringDeConexao := fmt.Sprintf("host=%s user=root password=root dbname=root port=5432 sslmode=disable", bdHost)
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
