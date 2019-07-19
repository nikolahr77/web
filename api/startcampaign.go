package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/web"
	"net/http"
)

func StartCampaign(cr web.CampaignRepository, ch chan web.Campaign) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		campaign, err := cr.Get(id)
		if err != nil {
			panic(err)
		}
		cr.SentStatus(id) //check for error

		ch <- campaign
		json.NewEncoder(w).Encode(http.StatusOK)
	}
}
