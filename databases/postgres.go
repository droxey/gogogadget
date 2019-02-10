package databases

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"

	"github.com/droxey/gogogadget/config"
)

func Init(conf *config.Config) *gorm.DB {
	connect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		conf.Postgres.Host,
		conf.Postgres.Port,
		conf.Postgres.User,
		conf.Postgres.Password,
		conf.Postgres.DB,
	)

	db, err := gorm.Open("postgres", connect)
	if err != nil {
		log.Panicln(err)
	}

	return db
}
