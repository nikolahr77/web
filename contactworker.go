package web

import (
	"database/sql"
	"fmt"
)

func GetContactInfo(campaign Campaign) error {
	connStr := "user=postgres dbname=mail sslmode=disable password=1234"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	query := `
	SELECT email FROM contacts WHERE age = $1 AND address = $2`
	var c Contact
	rows, err := db.Query(query, campaign.Segmentation.Age, campaign.Segmentation.Address)
	if err != nil {
		panic(err)
	}
	ContactSlice := make([]Contact, 1)
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&c.Email)
		if err != nil {
			return err
		}
		ContactSlice = append(ContactSlice, c)
	}

	ch := make(chan []Contact)
	go SendContact(ch, ContactSlice)
	fmt.Println(ContactSlice)
	//fmt.Println(campaign.Segmentation.Age)
	//fmt.Println(campaign.Segmentation.Address)
	//fmt.Println(c.Email)
	return nil
}

func ReceiveCampaignID(ch chan Campaign) {
	var cam Campaign
	for i := range ch {
		cam = i
	}
	GetContactInfo(cam)
}

func SendContact(ch chan []Contact, contact []Contact) {
	ch <- contact
	close(ch)
}
