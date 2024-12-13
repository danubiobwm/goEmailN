package endpoints

import (
	"net/http"
)

func (h *Handler) CampaingGet(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	campaign, err := h.CampaignService.Repository.Get()
	return campaign, 200, err
}
