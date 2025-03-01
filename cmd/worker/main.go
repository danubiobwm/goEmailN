package main

import (
	"log"
	"time"

	"github.com/danubiobwm/goEmailN/internal/domain/campaign"
	"github.com/danubiobwm/goEmailN/internal/infrastructure/database"
	"github.com/danubiobwm/goEmailN/internal/infrastructure/mail"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db := database.NewDb()
	repository := database.CampaignRepository{Db: db}

	campaignService := campaign.ServiceImp{
		Repository: &repository,
		SendMail:   mail.SendMail,
	}

	for {
		compaings, err := repository.GetCampaignToBeSent()

		if err != nil {
			println(err.Error())
		}
		println("amaount of campaigns", len(compaings))
		for _, campaign := range compaings {
			campaignService.SendEmailAndUpdateStatus(&campaign)
			println("campaign sent", campaign.ID)
		}
		time.Sleep(10 * time.Second)
	}
}
