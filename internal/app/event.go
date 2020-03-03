package app

import (
	"fmt"
	"regexp"

	"github.com/antonjah/leif/internal/pkg/constants"
	"github.com/antonjah/leif/internal/pkg/decide"
	"github.com/antonjah/leif/internal/pkg/flip"
	"github.com/antonjah/leif/internal/pkg/gitlab"
	"github.com/antonjah/leif/internal/pkg/questions"
	"github.com/antonjah/leif/internal/pkg/tldr"

	"github.com/nlopes/slack"
	"github.com/sirupsen/logrus"

	"github.com/antonjah/leif/internal/pkg/insults"
	"github.com/antonjah/leif/internal/pkg/lunches"
)

var (
	IGNORECASE = "(?i)"
	LUNCH      = regexp.MustCompile(IGNORECASE + "^\\.lunch (.*)")
	TACOS      = regexp.MustCompile(IGNORECASE + "^\\.tacos")
	QUESTION   = regexp.MustCompile(IGNORECASE + "^leif(.*)\\?")
	INSULT     = regexp.MustCompile(IGNORECASE + "^\\.insult (.*)")
	HELP       = regexp.MustCompile(IGNORECASE + "^\\.help")
	DECIDE     = regexp.MustCompile(IGNORECASE + "^\\.decide")
	FLIP       = regexp.MustCompile(IGNORECASE + "^\\.flip")
	TLDR       = regexp.MustCompile(IGNORECASE + "^\\.tldr (.*)")
	LOG        = regexp.MustCompile(IGNORECASE + "^\\.log (.*)")
	GITLAB     = regexp.MustCompile(IGNORECASE + "^\\.gitlab (.*)")
)

// EventHandler provides methods to handle slack events
type EventHandler struct {
	Client       *slack.Client
	LunchHandler lunches.LunchHandler
	Logger       *logrus.Logger
	Config       Config
}

// NewEventHandler returns a new EventHandler containing a slack client
// and LunchHandler
func NewEventHandler(client *slack.Client, lunchHandler lunches.LunchHandler, logger *logrus.Logger, config Config) EventHandler {
	return EventHandler{Client: client, LunchHandler: lunchHandler, Logger: logger, Config: config}
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
			e.Logger.Error(err)
			return
		}
		if len(matches) > 0 {
			HandleResponse(matches, event, e.Client, e.Logger)
			return
		}

		HandleResponse(fmt.Sprintf("Sorry, found nothing on %s", arg), event, e.Client, e.Logger)

	case TACOS.MatchString(event.Text):
		matches, err := e.LunchHandler.Search("taco")
		if err != nil {
			e.Logger.Error(err)
			return
		}
		if len(matches) > 0 {
			HandleResponse(matches, event, e.Client, e.Logger)
			return
		}

		HandleResponse("No restaurant is serving tacos today :white_frowning_face:", event, e.Client, e.Logger)

	case QUESTION.MatchString(event.Text):
		arg := QUESTION.FindStringSubmatch(event.Text)[1]
		answer := questions.GetAnswer(arg)
		HandleResponse(answer, event, e.Client, e.Logger)

	case INSULT.MatchString(event.Text):
		arg := INSULT.FindStringSubmatch(event.Text)[1]
		insult, err := insults.Get()
		if err != nil {
			e.Logger.Error(err)
		}
		HandleResponse(fmt.Sprintf("%s: %s", arg, insult), event, e.Client, e.Logger)

	case HELP.MatchString(event.Text):
		HandleResponse(constants.HELP, event, e.Client, e.Logger)

	case DECIDE.MatchString(event.Text):
		HandleResponse(decide.Get(), event, e.Client, e.Logger)

	case FLIP.MatchString(event.Text):
		HandleResponse(flip.Get(), event, e.Client, e.Logger)

	case TLDR.MatchString(event.Text):
		arg := TLDR.FindStringSubmatch(event.Text)[1]
		HandleResponse(tldr.GetTLDR(arg, e.Logger), event, e.Client, e.Logger)

	case LOG.MatchString(event.Text):
		arg := LOG.FindStringSubmatch(event.Text)[1]
		HandleResponse(GetLogRecords(arg), event, e.Client, e.Logger)

	case GITLAB.MatchString(event.Text):
		if e.Config.GitLabToken == "" || e.Config.GitLabURL == "" {
			HandleResponse("GitLab functionality is disabled, set ENVs to enable", event, e.Client, e.Logger)
			return
		}
		arg := GITLAB.FindStringSubmatch(event.Text)[1]
		HandleResponse(gitlab.Search(arg, e.Logger, e.Config.GitLabToken, e.Config.GitLabURL), event, e.Client, e.Logger)
	}
}

// TODO: Implement
func handleBotEvent() {}
