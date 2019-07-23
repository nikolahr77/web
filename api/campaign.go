package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/web"
	"github.com/web/convert"
	"net/http"
	"time"
)

//GetCampaign is used to get the ID from the GET request, sends a Get request
// and returns the Campaign with the same ID
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

//DeleteCampaign is used to delete a campaign with the ID from the DELETE request
func DeleteCampaign(cr web.CampaignRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		userID := context.Get(r, "userID").(string)
		err := cr.Delete(id, userID)
		if err != nil {
			http.Error(w, "Internal error", 500)
		}
	}
}

//CreateCampaign decodes JSON from the request and creates a new campaign based on the request
func CreateCampaign(cr web.CampaignRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c RequestCampaignDTO
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			http.Error(w, "Bad request", 400)
			return
		}
		cam := web.RequestCampaign{}
		userID := context.Get(r, "userID").(string)
		convert.SourceToDestination(c, &cam)
		campaign, err := cr.Create(cam, userID)
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

//UpdateCampaign selects a campaign with ID specified in the
// request and uses the JSON from the request to update the campaign
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
		userID := context.Get(r, "userID").(string)
		campaign, err := cr.Update(id, cam, userID)
		if err != nil {
			http.Error(w, "Internal error", 500)
			return
		}
		adaptedCam := CampaignDTO{}
		convert.SourceToDestination(campaign, &adaptedCam)
		json.NewEncoder(w).Encode(adaptedCam)
	}
}

//CampaignDTO is the campaign database object
type CampaignDTO struct {
	GUID         string          `json:"guid"`
	Name         string          `json:"name"`
	Segmentation SegmentationDTO `json:"segmentation"`
	Status       string          `json:"status"`
	CreatedOn    time.Time       `json:"created_on"`
	UpdatedOn    time.Time       `json:"updated_on"`
	MessageGUID  string          `json:"message_guid"`
}

//SegmentationDTO is part of the campaignDTO
type SegmentationDTO struct {
	Address string `json:"address"`
	Age     int    `json:"age"`
}

//RequestCampaignDTO is used to return info relevant to the user
type RequestCampaignDTO struct {
	Name         string          `json:"name"`
	Segmentation SegmentationDTO `json:"segmentation"`
	Status       string          `json:"status"`
	MessageGUID  string          `json:"message_guid"`
}
