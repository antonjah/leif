package jira

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"github.com/sirupsen/logrus"
	"net/http"
)

func GetIssue(arg string, logger *logrus.Logger, user, token, url string) string {
	auth := getAuth(user, token)
	client, err := jira.NewClient(auth.Client(), url)
	if err != nil {
		logger.Error(err)
		return "Failed to get jira issue, please check my logs"
	}

	issue, resp, err := client.Issue.Get(arg, nil)
	if resp != nil && resp.StatusCode == http.StatusNotFound {
		return fmt.Sprintf("Found no issue named '%s'", arg)
	}
	if err != nil {
		logger.Error(err)
		return "Failed to get jira issue, please check my logs"
	}

	return fmt.Sprintf("*%s:* *%+v*\n\n*Type:* %s\n*Priority:* %s\n\n%s", issue.Key, issue.Fields.Summary, issue.Fields.Type.Name, issue.Fields.Priority.Name, issue.Fields.Description)
}

func getAuth(user, token string) jira.BasicAuthTransport {
	return jira.BasicAuthTransport{
		Username: user,
		Password: token,
	}
}
