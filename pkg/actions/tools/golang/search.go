package golang

import (
	"net/url"
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
)

// ActionModuleSearch completes git repos in module format
//
//	github.com/carapace-sh/carapace@v0.0.1
//	github.com/spf13/cobra@v0.0.2
func ActionModuleSearch() carapace.Action {
	return carapace.ActionMultiParts("@", func(c carapace.Context) carapace.Action {
		switch len(c.Parts) {
		case 0:
			if strings.Count(c.Value, "/") < 3 {
				opts := git.SearchOpts{}.Default()
				opts.Prefix = false
				return git.ActionRepositorySearch(opts).NoSpace()
			}

			if strings.HasPrefix(c.Value, "github.com") {
				splitted := strings.Split(c.Value, "/")
				prefix := strings.Join(splitted[:3], "/") + "/"
				return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
					// TODO this should be done in ActionContents using `Context.Value` instead of `ContentOpts.Path`
					dir := filepath.Dir(c.Value)
					dirPrefix := dir + "/"
					if dir == "." {
						dir = ""
						dirPrefix = ""
					}
					return gh.ActionContents(gh.ContentOpts{
						Host:   splitted[0],
						Owner:  splitted[1],
						Name:   splitted[2],
						Branch: "",
						Path:   url.PathEscape(dir),
					}).Prefix(dirPrefix)
				}).Prefix(prefix)
			}

			return carapace.ActionValues()
		case 1:
			repo := c.Parts[0]
			if strings.HasPrefix(repo, "github.com") {
				if splitted := strings.SplitN(repo, "/", 4); len(splitted) > 3 {
					repo = strings.Join(splitted[:3], "/")
				}
			}
			return carapace.Batch(
				git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: "https://" + repo, Branches: true, Tags: true}),
				carapace.ActionValues("latest"),
			).ToA()
		default:
			return carapace.ActionValues()
		}
	})
}
