package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/web"
)

//StartCampaign decodes JSON from the request and starts a new campaign
func StartCampaign(cr web.CampaignRepository, ch chan web.Campaign) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		userID := r.Context().Value("userID").(string)
		campaign, err := cr.Get(id, userID)
		if err != nil {
			http.Error(w, "Internal error", 500)
			return
		}
		_, err = cr.SentStatus(id)
		if err != nil {
			http.Error(w, "Internal error", 500)
			return
		}
		ch <- campaign
		json.NewEncoder(w).Encode(http.StatusOK)
	}
}
