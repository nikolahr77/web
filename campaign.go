package web

type Segmentation struct {
	Address string
	Age     int
}

type Campaign struct {
	GUID         string
	Name         string
	Segmentation Segmentation
	Status       string
}

type RequestCampaign struct {
	Name         string
	Segmentation Segmentation
	Status       string
}

type CampaignRepository interface {
	Get(id string) (Campaign, error)
	Create(m RequestCampaign) (Campaign, error)
	Delete(id string) error
	Update(id string, m RequestCampaign) (Campaign, error)
}
