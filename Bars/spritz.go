package Bars

import (
	"strings"
	"time"

	"github.com/genesixx/coalibot/Struct"
	"github.com/genesixx/coalibot/Utils"
	"github.com/nlopes/slack"
)

func Spritz(option string, event *Struct.Message) bool {

	if strings.ToLower(strings.Split(option, " ")[0]) == "event" {
		eventBDE(event)
		return true
	}
	open := "Fermé !"
	color := "danger"
	if IsSpritzOpen() {
		open = "Ouvert !"
		color = "good"
	}
	params := Struct.SlackParams
	attachment := slack.Attachment{
		Color: color,
		Title: "Spritz",
		Text:  open,
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "Horaire",
				Value: "11-22h tous les jours sauf weekend et jours fériés!",
			},
			slack.AttachmentField{
				Title: "Ricard :sunny:",
				Value: "2.70€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Kronenbourg :beer:",
				Value: "Demi: 2.70€ Pinte: 5.00€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Carlsberg/Grimbergen :beer:",
				Value: "Demi: 3.70€ Pinte: 7.00€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "La Chouffe :beer: :gnome:",
				Value: "Demi: 3.70€ Pinte: 6.50€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Gallia :beer: :cucco:",
				Value: "Demi: 3.50€ Pinte: 6.00€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Liqueur/Shot :tumbler_glass:",
				Value: "2.00€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Cocktail " + Utils.Choice([]string{":cocktail2:", ":cocktail:"}),
				Value: "7.00€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Vin :wine_glass:",
				Value: "3.00€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Alcool fort :whiskey2:",
				Value: "6.00€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Suze :whiskey2:",
				Value: "3.00€",
				Short: true,
			},
		},
		Footer: "Powered by Coalibot",
	}
	params.Attachments = []slack.Attachment{attachment}
	event.API.PostMessage(event.Channel, "", params)
	return true
}

func IsSpritzOpen() bool {
	now := time.Now()
	if int(now.Weekday()) > 0 && int(now.Weekday()) < 6 && now.Hour() >= 11 && now.Hour() < 22 {
		return true
	}
	return false
}

func eventBDE(event *Struct.Message) {
	open := "Fermé !"
	color := "danger"
	now := time.Now()
	date, _ := time.Parse("2006-01-02", "2018-11-02")

	if now.Day() == date.Day() && now.Month() == date.Month() && now.Year() == date.Year() && now.Hour() >= 18 && now.Hour() < 21 || (now.Hour() == 21 && now.Minute() <= 30) {
		open = "Ouvert !"
		color = "good"
	}
	params := Struct.SlackParams
	attachment := slack.Attachment{
		Color: color,
		Title: "Spritz",
		Text:  open,
		Fields: []slack.AttachmentField{
			slack.AttachmentField{
				Title: "Event Blood Horror Party",
				Value: "Prix reduit sous presentation de billet de la soirée",
			},
			slack.AttachmentField{
				Title: "Ricard :sunny:",
				Value: "2.00€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Kronenbourg :beer:",
				Value: "Demi: 2.50€ Pinte: 4.00€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Gallia :beer: :cucco:",
				Value: "Demi: 3.00€ Pinte: 5.00€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Liqueur/Shot :tumbler_glass:",
				Value: "2.00€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Cocktail Special Halloween" + Utils.Choice([]string{":cocktail2:", ":cocktail:"}),
				Value: "5.00€",
				Short: true,
			},
			slack.AttachmentField{
				Title: "Whisky Coca / Vodka Pomme :whiskey2:",
				Value: "5.00€",
				Short: true,
			},
		},
		Footer: "Powered by Coalibot",
	}
	params.Attachments = []slack.Attachment{attachment}
	event.API.PostMessage(event.Channel, "", params)
}
