package web

type ContactWorker struct {
	ContactRepository ContactRepository
	contacts chan <- SendContacts //samo izprashta
	campaigns <- chan Campaign  //samo chete
	workers int
	stopChan chan struct{}
}

type SendContacts struct {
	MessageGUID string
	Contacts []Contact
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
	for i := 0; i < c.workers; i++ {

	}
}



func (c ContactWorker) GetContact() {
	for {
		select {
		case <- c.stopChan:
			return
		case campaign := <- c.campaigns:

			contacts, err := c.ContactRepository.GetAll(campaign.Segmentation)
			c.contacts <- contacts
		}
	}
}
