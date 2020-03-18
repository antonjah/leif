package app

import (
	"fmt"
	"regexp"

	"github.com/patrickmn/go-cache"

	"github.com/antonjah/leif/internal/pkg/constants"
	"github.com/antonjah/leif/internal/pkg/corona"
	"github.com/antonjah/leif/internal/pkg/decide"
	"github.com/antonjah/leif/internal/pkg/f1"
	"github.com/antonjah/leif/internal/pkg/flip"
	"github.com/antonjah/leif/internal/pkg/gitlab"
	"github.com/antonjah/leif/internal/pkg/postmord"
	"github.com/antonjah/leif/internal/pkg/questions"
	"github.com/antonjah/leif/internal/pkg/tldr"

	"github.com/nlopes/slack"
	"github.com/sirupsen/logrus"

	"github.com/antonjah/leif/internal/pkg/insults"
	"github.com/antonjah/leif/internal/pkg/lunches"
)

var (
	IGNORECASE = "(?i)"
	F1         = regexp.MustCompile(IGNORECASE + "^\\.f1")
	LUNCH      = regexp.MustCompile(IGNORECASE + "^\\.lunch (.*)")
	LUNCHES    = regexp.MustCompile(IGNORECASE + "^\\.lunches")
	TACOS      = regexp.MustCompile(IGNORECASE + "^\\.tacos")
	QUESTION   = regexp.MustCompile(IGNORECASE + "^leif(.*)\\?")
	INSULT     = regexp.MustCompile(IGNORECASE + "^\\.insult (.*)")
	HELP       = regexp.MustCompile(IGNORECASE + "^\\.help")
	DECIDE     = regexp.MustCompile(IGNORECASE + "^\\.decide")
	FLIP       = regexp.MustCompile(IGNORECASE + "^\\.flip")
	TLDR       = regexp.MustCompile(IGNORECASE + "^\\.tldr (.*)")
	LOG        = regexp.MustCompile(IGNORECASE + "^\\.log (.*)")
	GITLAB     = regexp.MustCompile(IGNORECASE + "^\\.gitlab (.*)")
	POSTMORD   = regexp.MustCompile(IGNORECASE + "^\\.postmord (.*)")
	CORONA     = regexp.MustCompile(IGNORECASE + "^\\.corona (.*)")
)

// EventHandler provides methods to handle slack events
type EventHandler struct {
	Client *slack.Client
	Cache  cache.Cache
	Logger *logrus.Logger
	Config Config
}

// NewEventHandler returns a new EventHandler containing a slack client
// and LunchHandler
func NewEventHandler(client *slack.Client, cache cache.Cache, logger *logrus.Logger, config Config) EventHandler {
	return EventHandler{Client: client, Cache: cache, Logger: logger, Config: config}
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
	case F1.MatchString(event.Text):
		HandleResponse(f1.GetLatestResult(e.Logger), event, e.Client, e.Logger)

	case LUNCHES.MatchString(event.Text):
		HandleResponse(lunches.GetAll(e.Cache, e.Logger), event, e.Client, e.Logger)

	case LUNCH.MatchString(event.Text):
		arg := LUNCH.FindStringSubmatch(event.Text)[1]
		matches, err := lunches.Search(arg, e.Cache, e.Logger)
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
		matches, err := lunches.Search("taco", e.Cache, e.Logger)
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

	case POSTMORD.MatchString(event.Text):
		if e.Config.PostMordToken == "" {
			HandleResponse("PostMord functionality is disabled, set ENVs to enable", event, e.Client, e.Logger)
			return
		}
		arg := POSTMORD.FindStringSubmatch(event.Text)[1]
		HandleResponse(postmord.Search(arg, e.Logger, e.Config.PostMordToken), event, e.Client, e.Logger)

	case CORONA.MatchString(event.Text):
		arg := CORONA.FindStringSubmatch(event.Text)[1]
		HandleResponse(corona.GetStatuses(arg, e.Cache, e.Logger), event, e.Client, e.Logger)
	}
}
