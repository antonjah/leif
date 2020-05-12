package constants

// HELP lists all commands known to Leif
const HELP = `*Random question*:
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
.corona <country> get the COVID19 status for a country`

// DagensLunchURL holds the URL to matochmat' daily lunch list
const DagensLunchURL = "https://www.matochmat.se/skelleftea"

// TLDRBaseURL holds the base used for formatting the GET TLDR URL
const TLDRBaseURL = "https://api.github.com/repos/tldr-pages/tldr/contents/pages/common/%s.md?ref=master"

// CoronaBaseURL holds the base for the Corona status API
const CoronaBaseURL = "https://api.covid19api.com/total/country/%s"

// InsultURL
const InsultURL = "https://evilinsult.com/generate_insult.php?lang=en&type=text"
