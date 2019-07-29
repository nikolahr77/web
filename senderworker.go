package web

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

//SenderWorker sends the payload
type SenderWorker struct {
	MessageRequests <-chan MessageRequest //samo chete
	Workers         int
	StopChan        chan struct{}
	ApiKey          string
	SecretKey       string
	SAPIHost        string
}

//Start starts a specified number of workers
func (s SenderWorker) Start() {
	for i := 0; i < s.Workers; i++ {
		go s.sendEmail()
	}
}

func (s SenderWorker) sendEmail() {
	for {
		select {
		case <-s.StopChan:
			return
		case mr := <-s.MessageRequests:
			byte, err := json.Marshal(mr)
			if err != nil {
				log.Print(err)
			}

			payload := strings.NewReader(string(byte))
			req, _ := http.NewRequest("POST", s.SAPIHost, payload)

			key := "Basic " + basicAuth(s.ApiKey, s.SecretKey)
			req.Header.Add("Content-Type", "application/json")
			req.Header.Add("Authorization", key)

			res, _ := http.DefaultClient.Do(req)

			defer res.Body.Close()
		}
	}
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
