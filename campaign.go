package web

import "time"

type Segmentation struct {
	GUID       string
	Address    string
	Age        int
	CampaignID string
}

type Campaign struct {
	GUID         string
	Name         string
	Segmentation Segmentation
	Status       string
	CreatedOn    time.Time
	UpdatedOn    time.Time
	MessageGUID  string
	UserID       string
}

type RequestCampaign struct {
	Name         string
	Segmentation Segmentation
	Status       string
	MessageGUID  string
	UserID       string
}

type RequestSegmentation struct {
	Address string
	Age     int
}

type CampaignRepository interface {
	Get(id string) (Campaign, error)
	Create(m RequestCampaign, userID string) (Campaign, error)
	Delete(id string) error
	Update(id string, m RequestCampaign) (Campaign, error)
	SentStatus(id string) (Campaign, error)
}
