package web_test

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/web"
	"github.com/web/api"
	"github.com/web/persistant"
	"log"
	"net/http"
	"testing"
)

func myhttphandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func startserver(t *testing.T, url string, want web.MessageRequest) {

	f := func(w http.ResponseWriter, r *http.Request) {
		got := web.MessageRequest{}

		json.NewDecoder(r.Body).Decode(&got)
		assert.Equal(t, got, want)
	}

	http.HandleFunc(url, myhttphandler)
	http.ListenAndServe(":8090", nil)
}

func TestSenderWorker_Start(t *testing.T) {
	connStr := "user=postgres dbname=mail sslmode=disable password=1234"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Print(err)
	}
	rc := persistant.RealClock{}
	clock := persistant.Clock(rc)

	cr := persistant.NewContactRepository(db, clock)
	msg := persistant.NewMessageRepository(db, clock)
	cam := persistant.NewCampaignRepository(db, clock)

	ch := make(chan web.Campaign)
	msgChan := make(chan web.MessageRequest)
	stopChan := make(chan struct{})

	api.StartCampaign(cam, ch)

	contactWorker := web.MessageRequestWorker{
		ContactRepository: cr,
		MessageRepository: msg,
		Campaigns:         ch,
		Messages:          msgChan,
		Workers:           2,
		StopChan:          stopChan,
		FromEmail:         "n.hristov@proxiad.com",
	}
	contactWorker.Start()

}

//Suzdavame worker (channetli reppository ....)
//Puskame FakeServer
//Suzdavame Expected
//Suzdavame msgRequest i go puskame po channel
//prashtam StopChan
