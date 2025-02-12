package endpoints

import "github.com/danubiobwm/goEmailN/internal/domain/campaign"

type Handler struct {
	CampaignService campaign.Service
}
