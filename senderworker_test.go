package web_test

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/web"
	"net/http"
	"testing"
)

func startserver(t *testing.T, url string, actual web.MessageRequest,serverChan chan interface{}) {
	r := mux.NewRouter()

	f := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("21111111")
		got := web.MessageRequest{}

		json.NewDecoder(r.Body).Decode(&got)
		assert.Equal(t, got, actual)
		fmt.Println(got)
		fmt.Println(actual)
		}
	r.HandleFunc(url, f).Methods("POST")
	serverChan <- nil
	http.ListenAndServe(":8090", r)
}

func TestSenderWorker_Start(t *testing.T) {
	Emails := web.Email{
		Email: "nikola@gmail.com",
		Name: "Nikola",
	}
	Sender := web.Email{
		Email: "Ivan@gmail.com",
		Name: "Ivan",
	}
	recipients := make([]web.Email,1)
	recipients[0] = Emails

	NewMessages := web.NewMessage{
		From: Sender,
		To: recipients,
		TextPart: "This is a test MSG",
	}

	TestMessageRequest := make([]web.NewMessage,1)
	TestMessageRequest[0] = NewMessages
	MSGRequest := web.MessageRequest{
		Messages: TestMessageRequest,
	}

	actual := web.MessageRequest{
		Messages: []web.NewMessage{NewMessages},
	}
	serverChan := make(chan interface{})

	go startserver(t,"/test", actual, serverChan)

	<- serverChan

	msgChan := make(chan web.MessageRequest)
	stopChan := make(chan struct{})


	senderWorker := web.SenderWorker{
		MessageRequests: msgChan,
		Workers:         1,
		StopChan:        stopChan,
		ApiKey:          "AK1234",
		SecretKey:       "SK4321",
		SAPIHost:        "http://localhost:8090/test",
	}

	senderWorker.Start()

	msgChan <- MSGRequest

	stopChan <- struct{}{}
}
