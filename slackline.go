package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/codegangsta/martini"
	//"github.com/nlopes/slack"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	//"strings"
)

const postMessageURL = "/services/hooks/incoming-webhook?token="
admins := [...]string{"thomasjosif", "sirius", "kigen", "homer", "imasonaz", "ruthless"}


type slackMessage struct {
	Channel   string `json:"channel"`
	Username  string `json:"username"`
	Text      string `json:"text"`
	//Avatar    string `json:"icon_url"`
	LinkNames bool   `json:"link_names"`
}

func (s slackMessage) payload() io.Reader {
	content := []byte("payload=")
	s.LinkNames = true
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

	if err != nil {
		return err
	}

	if res.StatusCode != 200 {
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		return errors.New(res.Status + " - " + string(body))
	}

	return
}

func stringInSlice(a string, list []string) bool {
    for _, b := range list {
        if b == a {
            return true
        }
    }
    return false
}

func main() {
    	//api := slack.New("")
	m := martini.Classic()
	m.Post("/bridge", func(res http.ResponseWriter, req *http.Request) {
		username := req.PostFormValue("user_name")
		text := req.PostFormValue("text")
		team := req.PostFormValue("team_domain")
		//userid := req.PostFormValue("user_id")

		if username == "slackbot" {
			// Avoid infinite loop
			return
		}
		if(stringInSlice(username, admins)) {
			editedusername := username + " (" + team + ") [Integrations Admin]"
		}
		else {
			editedusername := username + " (" + team + ") [Integrations Admin]"
		}
		
		// Get avatar.
		/*user, err := api.GetUserInfo(userid)
   		if err != nil {
	  	    fmt.Printf("%s\n", err)
	  	    return
   		}*/

		msg := slackMessage{
			Username: editedusername,
			Text:     text,
			//Avatar:   user.Profile.ImageOriginal,
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
