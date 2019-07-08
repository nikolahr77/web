package web

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type StartCamRepo struct {
	CampaignRepository CampaignRepository
}

func (cr StartCamRepo) StartCampaign() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		campaign, err := cr.CampaignRepository.Get(id)
		if err != nil {
			panic(err)
		}

		ch := make(chan Campaign)
		go ReceiveCampaignID(ch)
		SendCampaignID(ch, campaign)
		json.NewEncoder(w).Encode(http.StatusOK)
	}
}

func SendCampaignID(ch chan Campaign, campaign Campaign) {
	ch <- campaign
	close(ch)
}