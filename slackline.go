package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/codegangsta/martini"
	"github.com/nlopes/slack"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
)

const postMessageURL = "/services/hooks/incoming-webhook?token="

type slackMessage struct {
	Channel   string `json:"channel"`
	Username  string `json:"username"`
	Text      string `json:"text"`
	Avatar    string `json:"icon_url"`
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

	admins := map[string]bool {
    "thomasjosif": true,
    "sirius": true,
    "kigen": true,
    "ruthless": true,
    "imasonaz": true,
    "homer": true,
    "pyrrhic": true,
	}

	m := martini.Classic()
	m.Post("/bridge", func(res http.ResponseWriter, req *http.Request) {
		debug := false
		username := req.PostFormValue("user_name")
		text := req.PostFormValue("text")
		team := req.PostFormValue("team_domain")
		userid := req.PostFormValue("user_id")
		editedusername := "NULL"
		avatar := "nil"
		if username == "slackbot" {
			// Avoid infinite loop
			return
		}
		// Remove whitespace
		userid = strings.Replace(userid, " ", "", -1)

		if admins[username] {
			editedusername = username + " (" + team + ") [Admin]"
		} else {
			editedusername = username + " (" + team + ")"
		}

		if text == "!debug" {
			if(admins[username]){
				debug = true
				text = "This is a debugging message. Please ignore."
			}
		}


		if team == "hellsgamers" {
			hellsgamers := slack.New("")
			// Get avatar.
			hguser, hgerror := hellsgamers.GetUserInfo(userid)
	   		if hgerror != nil {
		  	    fmt.Printf("%s\n", hgerror)
		  	    return
	   		}
	   		avatar = hguser.Profile.ImageOriginal
		} else if team == "hg-ce" {
			hgce := slack.New("")
			// Get avatar.
			hgceuser, hgceerror := hgce.GetUserInfo(userid)
	   		if hgceerror != nil {
		  	    fmt.Printf("%s\n", hgceerror)
		  	    return
	   		}
	   		avatar = hgceuser.Profile.ImageOriginal
		} else if team == "hgdc" {
			hgdc := slack.New("")
			// Get avatar.
			hgdcuser, hgdcerror := hgdc.GetUserInfo(userid)
	   		if hgdcerror != nil {
		  	    fmt.Printf("%s\n", hgdcerror)
		  	    return
	   		}
	   		avatar = hgdcuser.Profile.ImageOriginal
		} else if team == "hgmods" {
			hgmods := slack.New("")
			// Get avatar.
			hgmodsuser, hgmodserror := hgmods.GetUserInfo(userid)
	   		if hgmodserror != nil {
		  	    fmt.Printf("%s\n", hgmodserror)
		  	    return
	   		}
	   		avatar = hgmodsuser.Profile.ImageOriginal
		} 

		if strings.Contains(text, "!announce") {
			if(admins[username]){
				text = strings.Replace(text, "!announce", "", -1)
				avatar = "http://icons.iconarchive.com/icons/graphicloads/100-flat/256/announcement-icon.png"
				editedusername = "Announcement"
			}
		}

		msg := slackMessage{
			Username: editedusername,
			Text:     text,
			Avatar:   avatar,
		}

		domain := req.URL.Query().Get("domain")
		token := req.URL.Query().Get("token")

		if debug {
			fmt.Printf("DEBUG: %s | %s | %s | %s | %s", username, editedusername, userid, team)
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
