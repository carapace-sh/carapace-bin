package action

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

// ActionTopicSearch completes topics by search
func ActionTopicSearch(cmd *cobra.Command) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionTopicSearch(gh.HostOpts{Host: repo.Host})
}

// ActionTopics completes topics of a user or organization
func ActionTopics(cmd *cobra.Command, user string) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionTopics(gh.TopicOpts{Host: repo.Host, User: user})
}

// ActionRepoTopics completes topics of a repository
func ActionRepoTopics(cmd *cobra.Command) carapace.Action {
	repo := repoOpts(cmd)
	return gh.ActionRepositoryTopics(gh.RepoOpts{Host: repo.Host, Owner: repo.Owner, Name: repo.Name})
}
