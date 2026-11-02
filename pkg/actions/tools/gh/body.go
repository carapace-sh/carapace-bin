package gh

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
)

type BodyOpts struct {
	Host  string
	Owner string
	Name  string
}

func (o BodyOpts) Default() BodyOpts {
	return o
}

func (o BodyOpts) repo() RepoOpts {
	return RepoOpts(o)
}

// ActionBody completes issue/pr bodies with keywords, mentions and emojis
//
//	closes #42
//	related #a/repo#12
//	:smile:
func ActionBody(opts BodyOpts) carapace.Action {
	return carapace.ActionMultiParts(" ", func(c carapace.Context) carapace.Action {
		switch {
		case strings.HasPrefix(c.Value, ":"):
			return ActionEmojis()

		case strings.HasPrefix(c.Value, "@"):
			c.Value = strings.TrimPrefix(c.Value, "@")
			return ActionMentionableUsers(opts.repo()).Invoke(c).Prefix("@").ToA()

		default:
			return actionKeywordLinks(opts)
		}
	})
}

func actionKeywordLinks(opts BodyOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		keywordsOfficial := []string{"close", "closes", "closed", "fix", "fixes", "fixed", "resolve", "resolves", "resolved"}
		keywordsCustom := []string{"causes", "caused-by", "closed_by", "fixed_by", "related", "resolved_by", "see"}

		re := regexp.MustCompile(fmt.Sprintf("^(%v)$", strings.Join(append(keywordsOfficial, keywordsCustom...), "|")))
		if last := len(c.Parts) - 1; last < 0 || !re.MatchString(c.Parts[last]) {
			return carapace.Batch(
				carapace.ActionValues(keywordsOfficial...).Style(styles.Gh.StateClosed),
				carapace.ActionValues(keywordsCustom...),
			).ToA().Invoke(c).Suffix(" #").ToA().NoSpace('#')
		}

		return carapace.ActionMultiParts("#", func(c carapace.Context) carapace.Action {
			switch len(c.Parts) {
			case 0:
				link := ActionOwnerRepositories(HostOpts{Host: opts.Host}).Invoke(c).Suffix("#").ToA()
				if strings.Contains(c.Value, "/") {
					return link
				}
				return ActionOwnerRepositories(HostOpts{Host: opts.Host})
			case 1:
				repo := opts.repo()
				if r := c.Parts[0]; r != "" {
					splitted := strings.SplitN(r, "/", 2)
					repo.Owner = splitted[0]
					if len(splitted) > 1 {
						repo.Name = splitted[1]
					}
				}
				return carapace.Batch(
					ActionIssues(IssueOpts{Host: repo.Host, Owner: repo.Owner, Name: repo.Name, Open: true, Closed: true}),
					ActionPullRequests(PullRequestOpts{Host: repo.Host, Owner: repo.Owner, Name: repo.Name, Open: true, Closed: true, Merged: true}),
				).ToA()
			default:
				return carapace.ActionValues()
			}
		})
	})
}
