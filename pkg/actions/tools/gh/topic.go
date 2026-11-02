package gh

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/carapace-sh/carapace"
)

type topic struct {
	Name             string
	ShortDescription string `json:"short_description"`
}

type topicQuery struct {
	Items []topic
}

type TopicOpts struct {
	Host string
	User string
}

func (o TopicOpts) Default() TopicOpts {
	return o
}

// ActionTopicSearch completes topics by search
//
//	dotnet (dotnet is the managed software framework from Microsoft)
//	golang (Go is a statically typed programming language)
func ActionTopicSearch(opts HostOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if len(c.Value) < 2 {
			return carapace.ActionMessage("topic search needs at least two characters")
		}

		var queryResult topicQuery
		return apiV3Action(opts.repo(), fmt.Sprintf(`search/topics?per_page=100&q=%v`, url.QueryEscape(c.Value)), &queryResult, func() carapace.Action {
			vals := make([]string, 0, len(queryResult.Items)*2)
			for _, topic := range queryResult.Items {
				vals = append(vals, topic.Name, topic.ShortDescription)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	}).Tag("topic search")
}

// ActionTopics completes topics of a user or organisation
//
//	cli
//	shell
//	completion
func ActionTopics(opts TopicOpts) carapace.Action {
	args := []string{"repo", "list", "--limit", "500", "--json", "repositoryTopics", "--jq", "[ .[].repositoryTopics | select(. != null) | .[].name ] | sort | unique | .[]"}
	if opts.User != "" {
		args = append(args, opts.User)
	}
	if opts.Host != "" {
		args = append(args, "--hostname", opts.Host)
	}
	return carapace.ActionExecCommand("gh", args...)(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		return carapace.ActionValues(lines[:len(lines)-1]...)
	})
}

type repoTopicsQuery struct {
	Names []string
}

// ActionRepoTopics completes topics of a repository
//
//	completion
//	shell
func ActionRepoTopics(opts RepoOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return resolveRepo(c, opts, func(resolved RepoOpts) carapace.Action {
			var queryResult repoTopicsQuery
			return apiV3Action(resolved, fmt.Sprintf("repos/%v/%v/topics", resolved.Owner, resolved.Name), &queryResult, func() carapace.Action {
				return carapace.ActionValues(queryResult.Names...)
			})
		})
	}).Tag("topics")
}
