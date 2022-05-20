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

const (
	HOST     = "db.host"
	PORT     = "db.port"
	USER     = "db.user"
	PASSWORD = "DB_PASSWORD"
	DBNAME   = "db.dbname"
	SSLMODE  = "db.sslmode"
)

func InitMemoryRepo() links.LinksRepo {
	return links.NewLinksMemoryRepo()
}

func InitSQLRepo() links.LinksRepo {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializating configs: %s", err.Error())
	}

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v",
		viper.GetString(USER),
		os.Getenv(PASSWORD),
		viper.GetString(HOST),
		viper.GetString(PORT),
		viper.GetString(DBNAME),
		viper.GetString(SSLMODE))

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		log.Fatalf("error opening db: %s", err.Error())
	}

	db.SetMaxOpenConns(1000)
	db.SetConnMaxLifetime(5 * time.Minute)

	err = db.Ping()
	if err != nil {
		log.Fatalf("pinging database error: %s", err.Error())
	}
	return links.NewLinksSQLRepo(db)
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("configs")
	return viper.ReadInConfig()
}
