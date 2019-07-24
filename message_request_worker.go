package web

import (
	"fmt"
	"log"
)

// MessageRequestWorker create the payload of the http request.
type MessageRequestWorker struct {
	ContactRepository ContactRepository
	MessageRepository MessageRepository
	Messages          chan<- MessageRequest
	Campaigns         <-chan Campaign //samo chete
	Workers           int
	StopChan          chan struct{}
	FromEmail         string
}

//Contacts contains a slice of the contacts with the desired specifications and the GUID
// of the message that needs to be sent
type Contacts struct {
	MessageGUID string
	Contacts    []Contact
}

//Start starts a specified number of workers
func (c MessageRequestWorker) Start() {
	for i := 0; i < c.Workers; i++ {
		go c.create()
	}
}

func (mrw MessageRequestWorker) create() {
	for {
		select {
		case <-mrw.StopChan:
			return
		case campaign := <-mrw.Campaigns:
			message, err := mrw.MessageRepository.Get(campaign.MessageGUID, campaign.UserID)
			if err != nil {
				log.Print(err)
			}
			contacts, err := mrw.ContactRepository.GetAll(campaign.Segmentation)
			if err != nil {
				log.Print(err)
			}

			fmt.Printf("%#v\n", contacts)
			fmt.Println("len", len(contacts))
			recipients := make([]Email, len(contacts))
			for i, c := range contacts {
				e := Email{Email: c.Email, Name: c.Name}
				recipients[i] = e
			}

			mrw.Messages <- MessageRequest{Messages: []NewMessage{{From: Email{Email: mrw.FromEmail}, To: recipients, TextPart: message.Content}}}
		}
	}
}

type MessageRequest struct {
	Messages []NewMessage
}

type NewMessage struct {
	From     Email   //az
	To       []Email //contacts
	TextPart string  //slagame content ot messages
}

type Email struct {
	Email string
	Name  string
}
