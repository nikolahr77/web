package api

import (
	"encoding/json"
	"github.com/web"
	"net/http"
	"time"
)


func CreateCampaign(cr web.CampaignRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c RequestCampaignDTO
		var s RequestSegmentationDTO
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			http.Error(w, "Internal error", 500)
		}
		err1 := json.NewDecoder(r.Body).Decode(&s)
		if err1 != nil {
			http.Error(w, "Internal error", 500)
		}
		cam := adaptToRequestCampaign
	}
}




type RequestCampaignDTO struct {
	GUID         string `json:"guid"`
	Name         string `json:"name"`
	Segmentation web.Segmentation `json:"segmentation"`
	Status       string `json:"status"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
}

type RequestSegmentationDTO struct {
	Address string `json:"address"`
	Age     int `json:"age"`
}

