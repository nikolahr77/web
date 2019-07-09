package web

func EmailSender(contacts []Contact) error{
	return nil
}




func ReceiveContacts(ch chan []Contact) {
	var ContactSlice []Contact
	for i := range ch {
		ContactSlice = i
	}
	EmailSender(ContactSlice)
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