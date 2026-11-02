package gh

import (
	"fmt"
	"github.com/carapace-sh/carapace"
	"strconv"
)

// ActionDeployKeyFields completes deploy key fields
//
//	id
//	key
func ActionDeployKeyFields() carapace.Action {
	return carapace.ActionValues(
		"createdAt",
		"id",
		"key",
		"readOnly",
		"title",
	)
}

type deployKey struct {
	Id    int
	Title string
	Key   string
}

// ActionDeployKeys completes deploy keys
//
//	123456 (changelog-key ssh-ed25519 ..m09M)
func ActionDeployKeys(opts RepoOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return resolveRepo(c, opts, func(resolved RepoOpts) carapace.Action {
			var queryResult []deployKey
			return apiV3Action(resolved, fmt.Sprintf("repos/%v/%v/keys", resolved.Owner, resolved.Name), &queryResult, func() carapace.Action {
				vals := make([]string, 0)
				for _, key := range queryResult {
					vals = append(vals, strconv.Itoa(key.Id), fmt.Sprintf("%v %v", key.Title, truncateMiddle(26, key.Key)))
				}
				return carapace.ActionValuesDescribed(vals...)
			})
		})
	}).Tag("deploy keys")
}

func truncateMiddle(maxWidth int, t string) string {
	if len(t) <= maxWidth {
		return t
	}

	ellipsis := "..."
	if maxWidth < len(ellipsis)+2 {
		return t[0:maxWidth]
	}

	halfWidth := (maxWidth - len(ellipsis)) / 2
	remainder := (maxWidth - len(ellipsis)) % 2
	return t[0:halfWidth+remainder] + ellipsis + t[len(t)-halfWidth:]
}
