package web

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)


func  StartCampaign(cr CampaignRepository, ch chan Campaign) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		campaign, err := cr.Get(id)
		if err != nil {
			panic(err)
		}

		ch := make(chan Campaign) //channel v main, poddava se tyk i v contactworker
		go ReceiveCampaignID(ch)
		SendCampaignID(ch, campaign)
		json.NewEncoder(w).Encode(http.StatusOK)
	}
}

func SendCampaignID(ch chan Campaign, campaign Campaign) {
	ch <- campaign
	close(ch)
}


//da otide v api packet