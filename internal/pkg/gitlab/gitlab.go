package gitlab

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/xanzy/go-gitlab"
)

type ResultSet struct {
	Projects      []*gitlab.Project
	MergeRequests []*gitlab.MergeRequest
	// Deprecated: Commits since introduction of ElasticSearch requirements
	Commits []*gitlab.Commit
}

func (r *ResultSet) ToString(arg string) string {
	var s string
	if len(r.MergeRequests) > 0 {
		s = "Merge requests:\n"
		for _, mr := range r.MergeRequests {
			s += fmt.Sprintf("[%s] %s: %s\n", mr.Author.Name, mr.Title, mr.WebURL)
		}
		s += "\n"
	}
	if len(r.Projects) > 0 {
		s += "Projects:\n"
		for _, proj := range r.Projects {
			s += fmt.Sprintf("[%s] %s: %s\n", proj.Name, proj.Description, proj.WebURL)
		}
		s += "\n"
	}
	if s == "" {
		return "Did not find anything on " + arg
	}

	return s
}

func Search(arg string, logger *logrus.Logger, token, url string) string {
	c := gitlab.NewClient(nil, token)
	c.SetBaseURL(url)

	opts := &gitlab.SearchOptions{PerPage: 20, Page: 1}

	var rs ResultSet

	mrs, r, err := c.Search.MergeRequests(arg, opts)
	if err != nil {
		logger.Error(err)
		return "Got an unexpected response while searching for merge requests, please check my logs"
	}
	if r.StatusCode/100 != 2 {
		logger.Error(r.Status)
		return "Got an unexpected response from GitLab, please check my logs"
	}
	rs.MergeRequests = mrs

	projs, r, err := c.Search.Projects(arg, opts)
	if err != nil {
		logger.Error(err)
		return "Got an unexpected response while searching for projects, please check my logs"
	}
	if r.StatusCode/100 != 2 {
		logger.Error(r.Status)
		return "Got an unexpected response from GitLab, please check my logs"
	}
	rs.Projects = projs

	return rs.ToString(arg)
}
