package action

import (
	"fmt"
	"net/url"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

type content struct {
	Name string
	Type string
}

func ActionContents(cmd *cobra.Command, path string, branch string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		repo, err := repoOverride(cmd)
		if err != nil {
			return carapace.ActionMessage(err.Error())
		}
		ref := ""
		if branch != "" {
			ref = fmt.Sprintf("ref=%v", branch)
		}

		var queryResult []content
		return ApiV3Action(cmd, fmt.Sprintf(`repos/%v/%v/contents/%v?%v`, repo.RepoOwner(), repo.RepoName(), url.PathEscape(path), ref), &queryResult, func() carapace.Action {
			vals := make([]string, 0)
			for _, content := range queryResult {
				name := content.Name
				if content.Type == "dir" {
					name = content.Name + "/"
				}
				vals = append(vals, name, style.ForPathExt(name))
			}
			return carapace.ActionStyledValues(vals...)
		})
	})
}
