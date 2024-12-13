package endpoints

import (
	"net/http"

	"github.com/danubiobwm/goEmailN/internal/contract"
	"github.com/go-chi/render"
)

func (h *Handler) CampaingPost(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
	var request contract.NewCampaign
	render.DecodeJSON(r.Body, &request)

	id, err := h.CampaignService.Create(request)

	return map[string]string{"id": id}, 201, err
}