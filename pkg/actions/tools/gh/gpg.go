package gh

import (
	"github.com/rsteube/carapace"
)

type gpgKey struct {
	KeyId string `json:"key_id"`
	Name  string
}

// ActionGpgKeys completes gpg keys
//
//	AABBCCDDEEFF112 (example)
//	AABBCCDDEEFF113 (another)
func ActionGpgKeys(opts HostOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var gpgKeys []gpgKey
		return apiV3Action(opts.repo(), "user/gpg_keys", &gpgKeys, func() carapace.Action {
			vals := make([]string, 0, len(gpgKeys))
			for _, key := range gpgKeys {
				vals = append(vals, key.KeyId, key.Name)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}
