package Citation

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"

	"github.com/genesixx/coalibot/Struct"
	"github.com/nlopes/slack"
)

func GLaDOS(option string, event *Struct.Message) bool {
	bytes, err := ioutil.ReadFile("Citation/glados.txt")

	if err != nil {
		fmt.Println(err)
		return false
	}
	splited := strings.Split(string(bytes), "\n")
	citation := splited[rand.Int()%len(splited)]
	params := slack.PostMessageParameters{UnfurlMedia: true, UnfurlLinks: true, Markdown: true, IconURL: "https://files.gamebanana.com/img/ss/srends/530-90_5a24c3979f637.jpg", Username: "GLaDOS"}
	event.API.PostMessage(event.Channel, "> "+citation, params)
	return true
}
