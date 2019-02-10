package databases

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	"github.com/droxey/gogogadget/config"
	"github.com/droxey/gogogadget/models"
)

// Init initializes the database and AutoMigrates the models.
func Init(conf *config.Config) *gorm.DB {
	connect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Postgres.Host,
		conf.Postgres.Port,
		conf.Postgres.User,
		conf.Postgres.Password,
		conf.Postgres.DB,
	)

	db, err := gorm.Open("postgres", connect)

	db.AutoMigrate(&models.Endpoint{}, &models.User{})

	if err != nil {
		log.Panicln(err)
	}

	return db
}
