package app

import (
	"fmt"
	"regexp"

	"github.com/antonjah/gleif/internal/constants"
	"github.com/antonjah/gleif/internal/decide"
	"github.com/antonjah/gleif/internal/flip"
	"github.com/antonjah/gleif/internal/questions"

	"github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"

	"github.com/antonjah/gleif/internal/insults"
	"github.com/antonjah/gleif/internal/lunches"
)

var (
	// IGNORECASE base for regex ignore case sensitive
	IGNORECASE = "(?i)"
	// LUNCH regex query matcher
	LUNCH = regexp.MustCompile(IGNORECASE + "^\\.lunch (.*)")
	// TACOS regex query matcher
	TACOS = regexp.MustCompile(IGNORECASE + "^\\.tacos")
	// QUESTION regex query matcher
	QUESTION = regexp.MustCompile(IGNORECASE + "^leif(.*)\\?")
	// INSULT regex query matcher
	INSULT = regexp.MustCompile(IGNORECASE + "^\\.insult (.*)")
	// HELP regex query matcher
	HELP = regexp.MustCompile(IGNORECASE + "^\\.help")
	// DECIDE regex query matcher
	DECIDE = regexp.MustCompile(IGNORECASE + "^\\.decide")
	// FLIP regex query matcher
	FLIP = regexp.MustCompile(IGNORECASE + "^\\.flip")
)

// EventHandler provides methods to handle slack events
type EventHandler struct {
	Client       *slack.Client
	LunchHandler lunches.LunchHandler
}

// NewEventHandler returns a new EventHandler containing a slack client
// and LunchHandler
func NewEventHandler(client *slack.Client, lunchHandler lunches.LunchHandler) EventHandler {
	return EventHandler{Client: client, LunchHandler: lunchHandler}
}

// Handle filters bot events and passes events to respective sub handler
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
			return
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
			return
		}
		if len(matches) > 0 {
			HandleResponse(matches, event, e.Client)
			return
		}

		HandleResponse("No restaurant is serving tacos today :white_frowning_face:", event, e.Client)

	case QUESTION.MatchString(event.Text):
		arg := QUESTION.FindStringSubmatch(event.Text)[1]
		answer := questions.GetAnswer(arg)
		HandleResponse(answer, event, e.Client)

	case INSULT.MatchString(event.Text):
		arg := INSULT.FindStringSubmatch(event.Text)[1]
		insult := insults.Get()
		HandleResponse(fmt.Sprintf("%s: %s", arg, insult), event, e.Client)

	case HELP.MatchString(event.Text):
		HandleResponse(constants.HELP, event, e.Client)

	case DECIDE.MatchString(event.Text):
		HandleResponse(decide.Get(), event, e.Client)

	case FLIP.MatchString(event.Text):
		HandleResponse(flip.Get(), event, e.Client)
	}
}

// TODO: Implement
func handleBotEvent() {}
