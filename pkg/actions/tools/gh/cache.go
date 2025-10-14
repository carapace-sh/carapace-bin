package gh

import (
	"fmt"
	"strconv"

	"github.com/carapace-sh/carapace"
)

type actionCache struct {
	ActionsCaches []struct {
		ID  int    `json:"id"`
		Ref string `json:"ref"`
	} `json:"actions_caches"`
}

// ActionCaches completes caches
//
//	41 (refs/pull/1718/merge)
//	42 (refs/heads/master)
func ActionCaches(opts RepoOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// TODO handle this generally for apiV3 actions
		if opts.Owner == "" {
			opts.Owner = ":owner"
		}
		if opts.Name == "" {
			opts.Name = ":repo"
		}
		var queryResult actionCache
		return apiV3Action(opts, fmt.Sprintf("repos/%v/%v/actions/caches?per_page=100&sort=last_accessed_at&direction=desc", opts.Owner, opts.Name), &queryResult, func() carapace.Action {
			vals := make([]string, 0)
			for _, cache := range queryResult.ActionsCaches {
				vals = append(vals, strconv.Itoa(cache.ID), cache.Ref)
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	}).Tag("caches")
}

// ActionCacheRefs completes cache refs
//
//	refs/pull/1718/merge
//	refs/heads/master
func ActionCacheRefs(opts RepoOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// TODO handle this generally for apiV3 actions
		if opts.Owner == "" {
			opts.Owner = ":owner"
		}
		if opts.Name == "" {
			opts.Name = ":repo"
		}
		var queryResult actionCache
		return apiV3Action(opts, fmt.Sprintf("repos/%v/%v/actions/caches?per_page=100&sort=last_accessed_at&direction=desc", opts.Owner, opts.Name), &queryResult, func() carapace.Action {
			unique := make(map[string]bool)
			for _, cache := range queryResult.ActionsCaches {
				unique[cache.Ref] = true
			}

			vals := make([]string, 0)
			for ref := range unique {
				vals = append(vals, ref)
			}
			return carapace.ActionValues(vals...)
		})
	}).Tag("cache refs")
}

// ActionCacheFields completes label fields
//
//	id
//	key
func ActionCacheFields() carapace.Action {
	return carapace.ActionValues(
		"createdAt",
		"id",
		"key",
		"lastAccessedAt",
		"ref",
		"sizeInBytes",
		"version",
	)
}
