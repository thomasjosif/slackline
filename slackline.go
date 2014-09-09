package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/codegangsta/martini"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
)

const postMessageURL = "/services/hooks/incoming-webhook?token="

type slackMessage struct {
	Channel  string `json:"channel"`
	Username string `json:"username"`
	Text     string `json:"text"`
}

func (s slackMessage) payload() io.Reader {
	content := []byte("payload=")
	json, _ := json.Marshal(s)
	content = append(content, json...)
	return bytes.NewReader(content)
}

var mentionRegexp = regexp.MustCompile("<@([^>]+)>")

func (s *slackMessage) containsMention() bool {
	return mentionRegexp.MatchString(s.Text)
}

func (s slackMessage) sendTo(domain, token string) (err error) {
	payload := s.payload()

	res, err := http.Post(
		"https://"+domain+postMessageURL+token,
		"application/x-www-form-urlencoded",
		payload,
	)

	if res.StatusCode != 200 {
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		return errors.New(res.Status + " - " + string(body))
	}

	return
}

func main() {
	m := martini.Classic()
	m.Post("/bridge", func(res http.ResponseWriter, req *http.Request) {
		username := req.PostFormValue("user_name")
		text := req.PostFormValue("text")

		if username == "slackbot" {
			// Avoid infinite loop
			return
		}

		msg := slackMessage{
			Username: username,
			Text:     text,
		}

		domain := req.URL.Query().Get("domain")
		token := req.URL.Query().Get("token")

		if os.Getenv("DEBUG_BRIDGE") == domain {
			fmt.Printf("Request: %v\n", req.PostForm)
			fmt.Printf("Message: %v\n", msg)
		}

		fmt.Printf("message=received domain=%s hasMention=%v token=%s\n", domain, msg.containsMention(), token)

		err := msg.sendTo(domain, token)

		if err != nil {
			fmt.Printf("message=error description=%#v\n", err.Error())
			res.WriteHeader(500)
		} else {
			fmt.Println("message=sent")
		}
	})
	m.Run()
}
