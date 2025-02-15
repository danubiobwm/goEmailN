package main

import (
	"github.com/danubiobwm/goEmailN/internal/domain/campaign"
	"github.com/danubiobwm/goEmailN/internal/endpoints"
	"github.com/danubiobwm/goEmailN/internal/infrastructure/database"
	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := chi.NewRouter()

	db := database.NewDb()

	campaignService := campaign.ServiceImp{
		Repository: &database.CampaignRepository{Db: db},
	}

	handler := endpoints.Handler{
		CampaignService: &campaignService,
	}

	r.Route("/campaigns", func(r chi.Router) {
		r.Use(endpoints.Auth)

		r.Post("/", endpoints.HandlerError(handler.CampaignPost))
		r.Get("/{id}", endpoints.HandlerError(handler.CampaignGetById))
		r.Delete("/delete/{id}", endpoints.HandlerError(handler.CampaignDelete))
	})
	http.ListenAndServe(":3000", r)
}
