// package git contains git related actions
package git

import (
	"fmt"
	"net/url"
	"path/filepath"

	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/carapace-sh/carapace/pkg/uid"
)

// Uid TODO experimental
func Uid(host string, opts ...string) func(s string, uc uid.Context) (*url.URL, error) {
	return func(s string, uc uid.Context) (*url.URL, error) {
		if length := len(opts); length%2 != 0 {
			return nil, fmt.Errorf("invalid amount of arguments [git.Uid]: %v", length)
		}

		gitDir, err := traverse.GitDir(uc)
		if err != nil {
			return nil, err
		}

		workTree, err := traverse.GitWorkTree(uc)
		if err != nil {
			return nil, err
		}

		uid := &url.URL{
			Scheme: "git",
			Host:   host,
			Path:   s,
		}
		values := uid.Query()
		values.Add("GIT_WORK_TREE", workTree)
		if rel, err := filepath.Rel(workTree, gitDir); err != nil || rel != ".git" {
			values.Add("GIT_DIR", gitDir)
		}
		for i := 0; i < len(opts); i += 2 {
			if opts[i+1] != "" { // implicitly skip empty values
				values.Add(opts[i], opts[i+1])
			}
		}
		uid.RawQuery = values.Encode()

		return uid, nil
	}
}
