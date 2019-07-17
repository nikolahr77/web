package web

import "fmt"

type SenderWorker struct {
	ContactRepository ContactRepository
	MessageRepository MessageRepository
	Contacts          <-chan SendContacts //samo izprashta
	Workers           int
	StopChan          chan struct{}
}

func (s SenderWorker) Start() {
	for i := 0; i < s.Workers; i++ {
		go s.SendEmail()
	}
}

func (s SenderWorker) SendEmail() {
	for {
		select {
		case <-s.StopChan:
			return
		case contacts := <-s.Contacts:
			message, err := s.MessageRepository.Get(contacts.MessageGUID)
			if err != nil {
				panic(err)
			}
			myEmail := MyEmailConstructor()
			receiverEmails := EmailConstructor(contacts)
			sendMessage := SendMessageConstructor(myEmail, receiverEmails, message)
			fmt.Println(sendMessage)
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

func EmailConstructor(contacts SendContacts) []Email {
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
