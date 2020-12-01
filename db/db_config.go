package db

import (
	"github.com/sirupsen/logrus"
	"github.com/teten-nugraha/mikro-backend/util"
	"log"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"github.com/teten-nugraha/mikro-backend/domain"
)

func InitDB(arg string) *gorm.DB {

	if(strings.Compare(util.PRODUCTION, arg) == 0) {
		err := godotenv.Load("production.env")
		if err != nil {
			log.Fatal("Error loading production.env file")
		}
		logrus.Info("Mikro Backend using Production DB Profile")
	}else{
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}
		logrus.Info("Mikro Backend using Development DB Profile")
	}

	err := godotenv.Load("prod.env")


	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")

	db, err := gorm.Open("mysql", dbUsername+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName)
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&domain.User{})

	return db
}
