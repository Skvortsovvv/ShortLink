package bootstrap

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"os"
	"testingTask/pkg/links"
	"time"
)

func InitMemoryRepo() links.LinksRepo {
	return links.NewLinksMemoryRepo()
}

func InitSQLRepo() links.LinksRepo {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializating configs: %s", err.Error())
	}

	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	psqlInfo := fmt.Sprintf(`host=%s port=%s user=%s password=%s dbname=%s sslmode=%s`,
		viper.GetString("db.host"),
		viper.GetString("db.port"),
		viper.GetString("db.user"),
		os.Getenv("DB_PASSWORD"),
		viper.GetString("db.dbname"),
		viper.GetString("db.sslmode"))

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalf("error opening db: %s", err.Error())
	}

	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(10 * time.Minute)

	err = db.Ping()
	if err != nil {
		log.Fatalf("pinging database error: %s", err.Error())
	}
	return links.NewLinksSQLRepo(db)
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../configs")
	return viper.ReadInConfig()
}
