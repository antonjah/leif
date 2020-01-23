package app

import (
	"fmt"
	"regexp"

	"github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"

	"github.com/antonjah/gleif/internal/lunches"
)

var (
	IGNORECASE = "(?i)"
	LUNCH      = regexp.MustCompile(IGNORECASE + "^\\.lunch (.*)")
	TACOS      = regexp.MustCompile(IGNORECASE + "^\\.tacos")
	QUESTION = regexp.MustCompile(IGNORECASE + "^leif(.*)\\?")
)

type EventHandler struct {
	Client       *slack.Client
	LunchHandler lunches.Handler
}

func NewEventHandler(client *slack.Client, lunchHandler lunches.Handler) EventHandler {
	return EventHandler{Client: client, LunchHandler: lunchHandler}
}

func (e EventHandler) Handle(msg slack.RTMEvent) {
	switch event := msg.Data.(type) {
	case *slack.MessageEvent:
		// Verify that the event is not a bot event
		// to prevent infinite loops
		if len(event.BotID) == 0 {
			go e.handleMessageEvent(event)
		}
	}
}

func (e EventHandler) handleMessageEvent(event *slack.MessageEvent) {
	switch {
	case LUNCH.MatchString(event.Text):
		arg := LUNCH.FindStringSubmatch(event.Text)[1]
		matches, err := e.LunchHandler.Search(arg)
		if err != nil {
			log.Error(err)
		}
		if len(matches) > 0 {
			HandleResponse(matches, event, e.Client)
			return
		}

		HandleResponse(fmt.Sprintf("Sorry, found nothing on %s", arg), event, e.Client)
	case TACOS.MatchString(event.Text):
		matches, err := e.LunchHandler.Search("taco")
		if err != nil {
			log.Error(err)
		}
		if len(matches) > 0 {
			HandleResponse(matches, event, e.Client)
			return
		}

		HandleResponse("No restaurant is serving tacos today :white_frowning_face:", event, e.Client)
	case QUESTION.MatchString(event.Text):
		arg := QUESTION.FindStringSubmatch(event.Text)[1]
		HandleResponse(fmt.Sprintf("Du kan va%s", arg), event, e.Client)
	}
}

// TODO: Implement
func handleBotEvent() {}
