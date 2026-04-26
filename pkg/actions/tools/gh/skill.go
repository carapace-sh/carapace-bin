package gh

import (
	"fmt"
	"net/url"

	"github.com/carapace-sh/carapace"
)

// ActionSkills completes skills
//
//	re-codebase-knowledge
//	add-educational-comments
func ActionSkills(opts RepoOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		var queryResult []content
		return apiV3Action(opts, fmt.Sprintf(`repos/%v/%v/contents/%v`, opts.Owner, opts.Name, url.PathEscape("skills")), &queryResult, func() carapace.Action {
			vals := make([]string, 0)
			for _, content := range queryResult {
				if content.Type != "dir" {
					continue
				}
				vals = append(vals, content.Name)
			}
			return carapace.ActionValues(vals...)
		})
	}).Tag("skills")

}
