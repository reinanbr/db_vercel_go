package psql_vercel


import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/jackc/pgx/v4/pgxpool"
)


func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Erro ao carregar o arquivo .env")
	}
}

// Conectar ao banco de dados
func ConnectDB() *pgxpool.Pool {
	LoadEnv()
	databaseURL := os.Getenv("DATABASE_URL")
	config, err := pgxpool.ParseConfig(databaseURL)
	if err != nil {
		log.Fatal("Erro ao analisar a URL do banco de dados: ", err)
	}

	conn, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal("Erro ao conectar ao banco de dados: ", err)
	}
	fmt.Println("Conexão com o banco de dados estabelecida.")
	return conn
}
