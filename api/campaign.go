package api

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/web"
	"net/http"
	"time"
)

func GetCampaign(cr web.CampaignRepository) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		campaign, err := cr.Get(id)
		if err != nil {
			http.Error(w, "Internal error", 500)
		}
		json.NewEncoder(w).Encode(adaptCamToDTO(campaign))
	}
}

func DeleteCampaign(cr web.CampaignRepository) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]
		err := cr.Delete(id)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Internal error", 500)
		}
	}
}

func CreateCampaign(cr web.CampaignRepository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var c RequestCampaignDTO
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			http.Error(w, "Internal error", 500)
		}
		cam := adaptToRequestCampaign(c)

		campaign, err := cr.Create(cam)
		if err != nil {
			http.Error(w, "Internal error", 500)
		}
		json.NewEncoder(w).Encode(adaptCamToDTO(campaign))
	}
}

func UpdateCampaign(cr web.CampaignRepository) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		var c RequestCampaignDTO
		err := json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			http.Error(w, "Internal error", 500)
		}
		cam := adaptToRequestCampaign(c)
		id := mux.Vars(r)["id"]
		campaign, err := cr.Update(id,cam)
		if err != nil {
			http.Error(w, "Internal error", 500)
		}
		json.NewEncoder(w).Encode(adaptCamToDTO(campaign))
	}
}

func adaptToRequestCampaign(c RequestCampaignDTO) web.RequestCampaign {
	return web.RequestCampaign{
		Name:         c.Name,
		Segmentation: adaptDTOtoSeg(c.Segmentation),
		Status:       c.Status,
	}
}

func adaptDTOtoSeg(c SegmentationDTO) web.Segmentation {
	return web.Segmentation{
		Address: c.Address,
		Age:     c.Age,
	}
}

func adaptCamToDTO(c web.Campaign) CampaignDTO {
	return CampaignDTO{
		GUID:         c.GUID,
		Name:         c.Name,
		Status:       c.Status,
		Segmentation: adaptSegToDTO(c.Segmentation),
		CreatedOn:    c.CreatedOn,
		UpdatedOn:    c.UpdatedOn,
	}
}

func adaptSegToDTO(c web.Segmentation) SegmentationDTO {
	return SegmentationDTO{
		Age:     c.Age,
		Address: c.Address,
	}
}

type CampaignDTO struct {
	GUID         string          `json:"guid"`
	Name         string          `json:"name"`
	Segmentation SegmentationDTO `json:"segmentation"`
	Status       string          `json:"status"`
	CreatedOn    time.Time       `json:"created_on"`
	UpdatedOn    time.Time       `json:"updated_on"`
}

type SegmentationDTO struct {
	Address string `json:"address"`
	Age     int    `json:"age"`
}

type RequestCampaignDTO struct {
	Name         string          `json:"name"`
	Segmentation SegmentationDTO `json:"segmentation"`
	Status       string          `json:"status"`
}
