package main

import (
	"./slack"
	"fmt"
	"os"
)

func main() {
	s := slack.InitSlack(os.Getenv("SLACK_URL"))
	// s.AddUsers("ansony", "ansony")
	s.Channel = "#random"
	s.Text = "関数作ってるのでテスト"
	s.Post()
	fmt.Println("exit")
}
