package gh

import (
	"fmt"

	"github.com/carapace-sh/carapace"
)

type hook struct {
	Id     int
	Name   string
	Active bool
	Config struct {
		Url string
	}
}

// ActionRepositoryHooks completes repository webhooks
//
//	1234567 (https://example.com/hook)
func ActionRepositoryHooks(opts RepoOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return resolveRepo(c, opts, func(resolved RepoOpts) carapace.Action {
			var queryResult []hook
			return apiV3Action(resolved, fmt.Sprintf("repos/%v/%v/hooks", resolved.Owner, resolved.Name), &queryResult, func() carapace.Action {
				return hookAction(queryResult)
			})
		})
	}).Tag("webhooks")
}

// ActionOrganizationHooks completes organization webhooks
//
//	1234567 (https://example.com/hook)
func ActionOrganizationHooks(opts OwnerOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult []hook
		return apiV3Action(opts.repo(), fmt.Sprintf("orgs/%v/hooks", opts.Owner), &queryResult, func() carapace.Action {
			return hookAction(queryResult)
		})
	}).Tag("webhooks")
}

func hookAction(hooks []hook) carapace.Action {
	vals := make([]string, 0, len(hooks)*2)
	for _, hook := range hooks {
		vals = append(vals, fmt.Sprintf("%v", hook.Id), fmt.Sprintf("%v (%v)", hook.Name, hook.Config.Url))
	}
	return carapace.ActionValuesDescribed(vals...)
}
