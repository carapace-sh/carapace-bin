package gh

import (
	"fmt"
	"net/url"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
)

type content struct {
	Name string
	Type string
}

type ContentOpts struct {
	Host   string
	Owner  string
	Name   string
	Branch string
	Path   string // TODO should be handled using `Context.Value`
}

func (c ContentOpts) repo() RepoOpts {
	return RepoOpts{Host: c.Host, Owner: c.Owner, Name: c.Name}
}

// ActionContents completes contents
//
//	/README.md
//	/.github/workflows/go.yml
func ActionContents(opts ContentOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		ref := ""
		if opts.Branch != "" {
			ref = fmt.Sprintf("ref=%v", opts.Branch)
		}

		var queryResult []content
		return apiV3Action(opts.repo(), fmt.Sprintf(`repos/%v/%v/contents/%v?%v`, opts.Owner, opts.Name, url.PathEscape(opts.Path), ref), &queryResult, func() carapace.Action {
			vals := make([]string, 0)
			for _, content := range queryResult {
				name := content.Name
				if content.Type == "dir" {
					name = content.Name + "/"
				}
				vals = append(vals, name, style.ForPathExt(name, c))
			}
			return carapace.ActionStyledValues(vals...).NoSpace('/')
		})
	}).Tag("contents")
}
