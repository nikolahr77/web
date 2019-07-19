package web

import (
	"fmt"
	"log"
)

type EmailWorker struct {
	ContactRepository ContactRepository
	MessageRepository MessageRepository
	Messages          chan<- MessageRequest
	Campaigns         <-chan Campaign //samo chete
	Workers           int
	StopChan          chan struct{}
}

type Contacts struct {
	MessageGUID string
	Contacts    []Contact
}

func (c EmailWorker) Start() {
	for i := 0; i < c.Workers; i++ {
		go c.GetContact()
	}
}

func (c EmailWorker) GetContact() {
	for {
		select {
		case <-c.StopChan:
			return
		case campaign := <-c.Campaigns:
			message, err := c.MessageRepository.Get(campaign.MessageGUID)
			if err != nil {
				log.Print(err)
			}
			contacts, err := c.ContactRepository.GetAll(campaign.Segmentation)
			if err != nil {
				log.Print(err)
			}
			SendContacts := Contacts{
				Contacts:    contacts,
				MessageGUID: campaign.MessageGUID,
			}

			myEmail := MyEmailConstructor()
			receiverEmails := EmailConstructor(SendContacts)
			sendMessage := SendMessageConstructor(myEmail, receiverEmails, message)
			fmt.Println(sendMessage)
			SendMessageSlice := make([]SendMessage, 1)
			SendMessageSlice = append(SendMessageSlice, sendMessage)
			MessageRequest := MessageRequest{
				Messages: SendMessageSlice,
			}
			c.Messages <- MessageRequest
		}
	}
}

func MyEmailConstructor() Email {
	return Email{
		email: "n.hristov@proxiad.com",
		name:  "Nikola Hristov",
	}
}

func SendMessageConstructor(myMail Email, receiverMails []Email, message Message) SendMessage {
	return SendMessage{
		From:     myMail,
		To:       receiverMails,
		TextPart: message.Content,
	}
}

func EmailConstructor(contacts Contacts) []Email {
	singleEmail := Email{}
	receiverEmails := make([]Email, 1)

	for _, x := range contacts.Contacts {
		singleEmail.name = x.Name
		singleEmail.email = x.Email
		receiverEmails = append(receiverEmails, singleEmail)
	}
	return receiverEmails
}

type MessageRequest struct {
	Messages []SendMessage
}

type SendMessage struct {
	From     Email   //az
	To       []Email //contacts
	TextPart string  //slagame content ot messages
}

type Email struct {
	email string
	name  string
}
