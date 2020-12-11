package db

import (
	"log"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/teten-nugraha/mikro-backend/domain"
	"github.com/teten-nugraha/mikro-backend/util"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

func InitDB(args []string) *gorm.DB {

	processENV(args)

	dbUsername := os.Getenv("DB_USERNAME")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")

	db, err := gorm.Open("mysql", dbUsername+":"+dbPassword+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?parseTime=true")
	if err != nil {
		panic(err)
	}

	//doMigrateDDL(db)

	return db
}

func doMigrateDDL(db *gorm.DB) {

	// Migrate if there are new file in domain
	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Kategori{})
	db.AutoMigrate(&domain.Merchant{})
}

func processENV(args []string) {

	if len(args) >= 1 {

		env := args[0]
		if strings.Compare(util.PRODUCTION, env) == 0 {
			err := godotenv.Load("production.env")
			if err != nil {
				log.Fatal("Error loading production.env file")
			}
			logrus.Info("Mikro Backend using Production DB Profile")
		} else if strings.Compare(util.DEVELOPMENT, env) == 0 {
			err := godotenv.Load(".env")
			if err != nil {
				log.Fatal("Error loading .env file")
			}
			logrus.Info("Mikro Backend using Development DB Profile")
		}
	}

}
