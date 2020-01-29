[![](https://img.shields.io/docker/pulls/antonjah/leif.svg)](https://microbadger.com/images/antonjah/leif "")
[![](https://images.microbadger.com/badges/image/antonjah/leif.svg)](https://microbadger.com/images/antonjah/leif "")
[![](https://images.microbadger.com/badges/version/antonjah/leif.svg)](https://microbadger.com/images/antonjah/leif "")

# <img src="img/leif.png" width="50" height="50"/> Leif

Leif is a really stupid Go bot for Slack

## Requirements

In order for Leif to run you need to:

1. Create a Slack [bot user](https://api.slack.com/bot-users#creating-bot-user).  
2. Then generate an OAuth2 bot token from the [apps page](https://api.slack.com/apps).  
The actual token can then be found under the Your Apps page if you click on the name of your created bot and then go to the `OAuth & Permissions` page and then `Bot User OAuth Access Token`.  
The token should start with xoxb-*.

## Running Leif

Make sure you have the generated tokens from the [requirements](#requirements) step and run:

```bash
docker run -e SLACK_TOKEN=<slack-token> antonjah/leif
```

## Commands

List of possible commands

### Asking a question

```text
Leif <question>?
```

Examples:

```text
Leif what is the meaning of life?
Leif what is the weather like in Skellefteå?
Leif how high is the empire state building?
Leif what's the IMDB score for Band of Brothers?
```

### Lunch ([norran](https://norran.se))

Search for something you'd like to eat or get a specific menu for a restaurant

```bash
.lunch carbonara
.lunch pizzeria mama mia
```

### Tacos

If you want to know if any restaurants are serving tacos today, there's a command for that!

```bash
.tacos
```

### Insult

Insulting someone. What more do you need?

```bash
.insult @Simon
```

### TLDR

Leif will try to get information on how to use a specific command

```bash
.tldr docker-compose
.tldr grep
```

### Flip

Flip a coin

```bash
.flip
```

### Decide

Make Leif decide something for you

```bash
.decide do we really need these unit tests?
```

### Help

Output help and known commands

```bash
.help
```
