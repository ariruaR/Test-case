package configReader

import (
	"log"
	"os"
	"strconv"

	env "github.com/joho/godotenv"
)

type Config struct {
	Host     string `json:"psqlHost"`
	Port     int    `json:"psqlPort"`
	User     string `json:"psqlUser"`
	Password string `json:"psqlPassword"`
	DBname   string `json:"psqlDBname"`
}

func init() {
	if err := env.Load(); err != nil {
		log.Println(err)
	}
}
func ReadConfig() Config {
	psqlHost, _ := os.LookupEnv("PSQL_HOST")
	psqlPort, _ := os.LookupEnv("PSQL_PORT")
	psqlUser, _ := os.LookupEnv("PSQL_USER")
	psqlPassword, _ := os.LookupEnv("PSQL_PASSWORD")
	psqlDBname, _ := os.LookupEnv("PSQL_DBNAME")

	port, _ := strconv.Atoi(psqlPort)

	return Config{
		Host:     psqlHost,
		Port:     port,
		User:     psqlUser,
		Password: psqlPassword,
		DBname:   psqlDBname,
	}
}
