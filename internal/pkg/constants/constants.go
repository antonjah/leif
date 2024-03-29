package constants

const (
	DagensLunchURL   = "https://www.matochmat.se/lunch/skelleftea"
	TLDRBaseURL      = "https://api.github.com/repos/tldr-pages/tldr/contents/pages/common/%s.md?ref=master"
	CoronaBaseURL    = "https://api.covid19api.com/total/country/%s"
	InsultURL        = "https://evilinsult.com/generate_insult.php?lang=en&type=text"
	RandomWordApiURL = "https://random-word-api.herokuapp.com/word?number=2&swear=1"
	HELP             = `*Random question*:
Leif <question>?
*Specific commands*:
.f1 - Get the latest grand prix results
.lunch <restaurant> - Show the lunch menu for given restaurant
.lunches - Show a list of all lunches
.gitlab <search> - Search GitLab commits, projects and merge requests
.tacos - Find out if a nearby restaurant is serving some kind of taco today
.insult <user> - Make Leif insult given username
.help - Show this help section
.decide <arg> - Decide whether or not given argument is true or false
.flip - Flip a coin
.tldr <cmd> - Find information about a command
.log <level> - Get Leif logs
.postmord <parcel-id> get the shipment status for your parcel(s) from PostNord
.corona <country> get the COVID19 status for a country
.suggest - get a nice word to be used and abused
.jira <arg> - Get information about a JIRA issue`
)
