package main

import (
	"net/http"

	"github.com/danubiobwm/goEmailN/internal/domain/campaign"
	"github.com/danubiobwm/goEmailN/internal/endpoints"
	"github.com/danubiobwm/goEmailN/internal/infrastructure/database"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	campaignService := campaign.Service{
		Repository: &database.CampaignRepository{},
	}

	handler := endpoints.Handler{
		CampaignService: campaignService,
	}

	r.Post("/campaigns", handler.CampaingPost)
	r.Get("/campaigns", handler.CampaingGet)

	http.ListenAndServe(":3000", r)

}
