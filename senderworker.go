package web

type SenderWorker struct {
	ContactRepository ContactRepository
	MessageRepository MessageRepository
	Contacts          <-chan SendContacts //samo izprashta
	Workers           int
	StopChan          chan struct{}
}

func (s SenderWorker) Start() {
	for i := 0; i < s.Workers; i++ {
		//go c.GetContact()
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

			var SendMessage struct{}
			for i := range contacts.Contacts {
				contacts.Contacts[i]
			}
		}
	}
}

func SendMessageConstructor(contacts SendContacts, message Message) SendMessage {
	return SendMessage{
		From: "n.hristov@proxiad.com",
		To:   contacts,
	}
}

//func ReceiveContacts(ch chan []Contact) {
//	var ContactSlice []Contact
//	for i := range ch {
//		ContactSlice = i
//	}
//	EmailSender(ContactSlice)
//}
//
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
