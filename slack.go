package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type Slack struct {
	Url      string
	Channel  string
	Username string
	Text     string
	Users    []string
}

func InitSlack(url string) *Slack {
	slack := &Slack{}
	slack.Url = url
	return slack
}

func (s *Slack) AddUsers(users ...string) {
	s.Users = append(s.Users, users...)
}

func (s *Slack) Post() {
	if s.Text == "" {
		return
	}

	text := ""
	if len(s.Users) != 0 {
		text += "<@" + strings.Join(s.Users, "> <@") + "> "
	}
	text += s.Text

	postData := map[string]string{
		"text": text,
	}
	if s.Username != "" {
		postData["username"] = s.Username
	}
	if s.Channel != "" && strings.HasPrefix(s.Channel, "#") {
		postData["channel"] = s.Channel
	}
	postJson, err := json.Marshal(postData)
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest(http.MethodPost, s.Url, bytes.NewBuffer(postJson))
	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	if res.StatusCode != http.StatusOK {
		fmt.Println("Status Error! status: " + res.Status)
	}

	defer res.Body.Close()
}
