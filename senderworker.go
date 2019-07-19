package web

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type SenderWorker struct {
	MessageRequests <-chan MessageRequest //samo chete
	Workers         int
	StopChan        chan struct{}
	ApiKey          string
	SecretKey       string
	SAPIHost        string
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
			body, _ := ioutil.ReadAll(res.Body)

			fmt.Println(string(body))
		}
	}
}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}
