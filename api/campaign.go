package api

import (
	"encoding/json"
	"fmt"
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
			fmt.Printf("Error  %e", err)
			http.Error(w, "Internal error", 500)
		}
		fmt.Println(c)
		err1 := json.NewDecoder(r.Body).Decode(&s)
		if err1 != nil {
			fmt.Printf("Error 1 %e", err1)
			http.Error(w, "Internal error", 500)
		}
		fmt.Println(s)
		cam := adaptToRequestCampaign(c)
		seg := adaptToRequestSegmentation(s)

		campaign, segmentation, err2 := cr.Create(cam,seg)
		if err2 != nil {
			fmt.Printf("Error 2 %e", err2)
			http.Error(w, "Internal error", 500)
		}
		json.NewEncoder(w).Encode(adaptCamToDTO(campaign))
		json.NewEncoder(w).Encode(adaptSegToDTO(segmentation))
	}
}

func adaptToRequestCampaign(c RequestCampaignDTO) web.RequestCampaign{
	return web.RequestCampaign{
		Name: c.Name,
		Segmentation: c.Segmentation,
		Status: c.Status,
	}
}

func adaptToRequestSegmentation(s RequestSegmentationDTO) web.RequestSegmentation{
	return web.RequestSegmentation{
		Address: s.Address,
		Age: s.Age,
	}
}

func adaptCamToDTO (c web.Campaign) CampaignDTO{
	return CampaignDTO{
		GUID: c.GUID,
		Name: c.Name,
		Status: c.Status,
		Segmentation: c.Segmentation,
		CreatedOn: c.CreatedOn,
		UpdatedOn: c.UpdatedOn,
	}
}

func adaptSegToDTO (c web.Segmentation) SegmentationDTO{
	return SegmentationDTO{
		GUID: c.GUID,
		Age: c.Age,
		Address: c.Address,
		CampaignID: c.CampaignID,
	}
}

type CampaignDTO struct {
	GUID         string `json:"guid"`
	Name         string `json:"name"`
	Segmentation web.Segmentation `json:"segmentation"`
	Status       string `json:"status"`
	CreatedOn time.Time `json:"created_on"`
	UpdatedOn time.Time `json:"updated_on"`
}

type SegmentationDTO struct {
	GUID string `json:"guid"`
	CampaignID string `json:"campaign_id"`
	Address string `json:"address"`
	Age     int `json:"age"`
}

type RequestCampaignDTO struct {
	Name         string `json:"name"`
	Segmentation web.Segmentation `json:"segmentation"`
	Status       string `json:"status"`
}

type RequestSegmentationDTO struct {
	Address string `json:"address"`
	Age     int `json:"age"`
}

