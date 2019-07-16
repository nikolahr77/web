package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/web"
	"github.com/web/convert"
	"net/http"
	"time"
)

func GetCampaign(cr web.CampaignRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		campaign, err := cr.Get(id)
		if err != nil {
			http.Error(w, "Internal error", 500)
			return
		}
		adaptedCam := CampaignDTO{}
		convert.SourceToDestination(campaign, &adaptedCam)
		json.NewEncoder(w).Encode(adaptedCam)
	}
}

func DeleteCampaign(cr web.CampaignRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		err := cr.Delete(id)
		if err != nil {
			http.Error(w, "Internal error", 500)
		}
	}
}

func CreateCampaign(cr web.CampaignRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c RequestCampaignDTO
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			http.Error(w, "Bad request", 400)
			return
		}
		cam := web.RequestCampaign{}
		convert.SourceToDestination(c, &cam)
		campaign, err := cr.Create(cam)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal error", 500)
			return
		}
		camDTO := CampaignDTO{}
		convert.SourceToDestination(campaign, &camDTO)
		json.NewEncoder(w).Encode(camDTO)
	}
}

func UpdateCampaign(cr web.CampaignRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c RequestCampaignDTO
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			http.Error(w, "Bad request", 400)
			return
		}
		cam := web.RequestCampaign{}
		convert.SourceToDestination(c, &cam)
		id := mux.Vars(r)["id"]
		campaign, err := cr.Update(id, cam)
		if err != nil {
			http.Error(w, "Internal error", 500)
			return
		}
		adaptedCam := CampaignDTO{}
		convert.SourceToDestination(campaign, &adaptedCam)
		json.NewEncoder(w).Encode(adaptedCam)
	}
}

type CampaignDTO struct {
	GUID         string          `json:"guid"`
	Name         string          `json:"name"`
	Segmentation SegmentationDTO `json:"segmentation"`
	Status       string          `json:"status"`
	CreatedOn    time.Time       `json:"created_on"`
	UpdatedOn    time.Time       `json:"updated_on"`
	MessageGUID  string          `json:"message_guid"`
}

type SegmentationDTO struct {
	Address string `json:"address"`
	Age     int    `json:"age"`
}

type RequestCampaignDTO struct {
	Name         string          `json:"name"`
	Segmentation SegmentationDTO `json:"segmentation"`
	Status       string          `json:"status"`
	MessageGUID  string          `json:"message_guid"`
}
