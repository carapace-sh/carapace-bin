package gh

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

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

	// Path lists the contents of this directory instead of deriving the
	// current directory from Context.Value (i.e. the one being completed).
	Path string
}

func (c ContentOpts) repo() RepoOpts {
	return RepoOpts{Host: c.Host, Owner: c.Owner, Name: c.Name}
}

// ActionReadDirFields completes repo read-dir fields
//
//	gitSHA
//	gitType
func ActionReadDirFields() carapace.Action {
	return carapace.ActionValues(
		"gitSHA",
		"gitType",
		"mode",
		"modeOctal",
		"name",
		"nameRaw",
		"path",
		"pathRaw",
		"size",
		"submodule",
		"type",
	)
}

// ActionReadFileFields completes repo read-file fields
//
//	content
//	downloadUrl
func ActionReadFileFields() carapace.Action {
	return carapace.ActionValues(
		"content",
		"downloadUrl",
		"encoding",
		"gitSHA",
		"gitUrl",
		"htmlUrl",
		"name",
		"path",
		"size",
		"type",
		"url",
	)
}

// ActionContents completes contents
//
//	/README.md
//	/.github/workflows/go.yml
//
// ActionContents completes contents
//
//	/.github/
//	/README.md
func ActionContents(opts ContentOpts) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		path := opts.Path
		if path == "" {
			path = strings.TrimPrefix(filepath.Dir(c.Value), "/")
		}
		return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
			ref := ""
			if opts.Branch != "" {
				ref = fmt.Sprintf("ref=%v", opts.Branch)
			}

			return resolveRepo(c, opts.repo(), func(resolved RepoOpts) carapace.Action {
				var queryResult []content
				return apiV3Action(resolved, fmt.Sprintf(`repos/%v/%v/contents/%v?%v`, resolved.Owner, resolved.Name, url.PathEscape(path), ref), &queryResult, func() carapace.Action {
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
			})
		})
	}).Tag("contents")
}
