package web

type ContactWorker struct {
	ContactRepository ContactRepository
	Contacts          chan<- SendContacts //samo izprashta
	Campaigns         <-chan Campaign     //samo chete
	Workers           int
	StopChan          chan struct{}
}

type SendContacts struct {
	MessageGUID string
	Contacts    []Contact
}

//func NewContactWorker(repository ContactRepository, contacts chan []Contact, campaigns chan Campaign, workers int) ContactWorker{
//	return ContactWorker{
//		ContactRepository: repository,
//		contacts: contacts,
//		campaigns: campaigns,
//		workers: workers,
//	}
//}

func (c ContactWorker) Start() {
	for i := 0; i < c.Workers; i++ {
		go c.GetContact()
	}
}

func (c ContactWorker) GetContact() {
	for {
		select {
		case <-c.StopChan:
			return
		case campaign := <-c.Campaigns:
			contacts, err := c.ContactRepository.GetAll(campaign.Segmentation)
			if err != nil {
				panic(err)

			}
			SendContacts := SendContactsConstructor(contacts, campaign.MessageGUID)

			c.Contacts <- SendContacts
		}
	}
}

func SendContactsConstructor(contacts []Contact, messageID string) SendContacts {
	return SendContacts{
		Contacts:    contacts,
		MessageGUID: messageID,
	}
}
