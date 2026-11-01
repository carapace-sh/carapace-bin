package hg

import (
	"fmt"
	"net/url"

	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/carapace-sh/carapace/pkg/uid"
)

// Uid TODO experimental
func Uid(host string, opts ...string) func(s string, uc uid.Context) (*url.URL, error) {
	return func(s string, uc uid.Context) (*url.URL, error) {
		if length := len(opts); length%2 != 0 {
			return nil, fmt.Errorf("invalid amount of arguments [hg.Uid]: %v", length)
		}

		uid := &url.URL{
			Scheme: "hg",
			Host:   host,
			Path:   s,
		}
		values := uid.Query()
		if workTree, err := traverse.Parent(".hg")(uc); err == nil {
			values.Add("HG_WORK_TREE", workTree) // hg has no environment variable for the repository location
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
