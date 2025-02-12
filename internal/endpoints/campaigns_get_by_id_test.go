package endpoints

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/danubiobwm/goEmailN/internal/contract"
	internalmock "github.com/danubiobwm/goEmailN/internal/test/internal-mock"
	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_CampaignsGetById_should_return_campaign(t *testing.T) {
	assert := assert.New(t)
	campaign := contract.CampaignResponse{
		ID:      "343",
		Name:    "Test",
		Content: "Hi!",
		Status:  "Pending",
	}
	service := new(internalmock.CampaignServiceMock)
	service.On("GetBy", "343").Return(&campaign, nil)

	handler := Handler{CampaignService: service}
	req := httptest.NewRequest("GET", "/campaigns/343", nil)
	// Adiciona o contexto de rota com o parâmetro "id"
	ctx := chi.NewRouteContext()
	ctx.URLParams.Add("id", "343")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, ctx))

	rr := httptest.NewRecorder()

	response, status, _ := handler.CampaignGetById(rr, req)

	assert.Equal(200, status)
	assert.Equal(campaign.ID, response.(*contract.CampaignResponse).ID)
	assert.Equal(campaign.Name, response.(*contract.CampaignResponse).Name)
}

func Test_CampaignsGetById_should_return_error_when_something_wrong(t *testing.T) {
	assert := assert.New(t)
	service := new(internalmock.CampaignServiceMock)
	errExpected := errors.New("something wrong")
	service.On("GetBy", mock.Anything).Return(nil, errExpected)
	handler := Handler{CampaignService: service}
	req, _ := http.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	_, _, errReturned := handler.CampaignGetById(rr, req)

	assert.Equal(errExpected.Error(), errReturned.Error())
}
