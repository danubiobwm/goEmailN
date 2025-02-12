package database

import (
	"github.com/danubiobwm/goEmailN/internal/domain/campaign"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {
	dsn := "host=localhost user=emailn_dev password=Danu1985 dbname=emailn_dev port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})
	return db
}
