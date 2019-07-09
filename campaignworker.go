package web

import (
	"database/sql"
	"fmt"
)

type ContactInfo struct {
	email string
}

func GetContactInfo(campaign Campaign) error {
	connStr := "user=postgres dbname=mail sslmode=disable password=1234"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	query := `
	SELECT email FROM contacts WHERE age = $1 AND address = $2`
	var c ContactInfo
	rows, err := db.Query(query, campaign.Segmentation.Age, campaign.Segmentation.Address)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&c.email)
		if err != nil {
			return err
		}
	}
	fmt.Println(campaign.Segmentation.Age)
	fmt.Println(campaign.Segmentation.Address)
	fmt.Println(c.email)
	return nil
}

func ReceiveCampaignID(ch chan Campaign) {
	var cam Campaign
	for i := range ch {
		cam = i
	}
	GetContactInfo(cam)
}
