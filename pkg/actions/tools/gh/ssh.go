package gh

import (
	"strconv"

	"github.com/carapace-sh/carapace"
)

type sshKey struct {
	Id    int
	Title string
}

// ActionSshKeys completes ssh keys
//
//	12345678 (title)
//	23456789 (another)
func ActionSshKeys(opts HostOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var sshKeys []sshKey
		return apiV3Action(opts.repo(), "user/keys", &sshKeys, func() carapace.Action {
			vals := make([]string, 0, len(sshKeys))
			for _, key := range sshKeys {
				vals = append(vals, strconv.Itoa(key.Id), key.Title)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
