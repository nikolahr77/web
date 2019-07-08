package web

import "fmt"

type CampaignRepo struct {
	CampaignRepository CampaignRepository
}


func GetContactInfo(campaign Campaign) error{


	fmt.Println(campaign.Segmentation.Age)
	fmt.Println(campaign.Segmentation.Address)
	return nil
}


func ReceiveCampaignID(ch chan Campaign) {
	var cam Campaign
	for i := range ch{
		cam = i
	}
	GetContactInfo(cam)
}
