package gh

import (
	"fmt"
	"strconv"

	"github.com/carapace-sh/carapace"
)

// ActionAutolinkFields completes autolink fields
//
//	id
//	isAlphanumeric
func ActionAutolinkFields() carapace.Action {
	return carapace.ActionValues(
		"id",
		"isAlphanumeric",
		"keyPrefix",
		"urlTemplate",
	)
}

type autolink struct {
	ID        int    `json:"id"`
	KeyPrefix string `json:"key_prefix"`
}

// ActionAutolinks completes autolinks
//
//	1111111 (TICKET-)
//	2222222 (STORY-)
func ActionAutolinks(opts RepoOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return resolveRepo(c, opts, func(resolved RepoOpts) carapace.Action {
			var queryResult []autolink
			return apiV3Action(resolved, fmt.Sprintf("repos/%v/%v/autolinks", resolved.Owner, resolved.Name), &queryResult, func() carapace.Action {
				vals := make([]string, 0)
				for _, al := range queryResult {
					vals = append(vals, strconv.Itoa(al.ID), al.KeyPrefix)
				}
				return carapace.ActionValuesDescribed(vals...)
			})
		})
	}).Tag("auto links")
}
