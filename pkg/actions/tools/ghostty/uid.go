package ghostty

import (
	"fmt"
	"net/url"

	"github.com/carapace-sh/carapace/pkg/uid"
)

func Uid(host string, opts ...string) func(s string, uc uid.Context) (*url.URL, error) {
	return func(s string, uc uid.Context) (*url.URL, error) {
		if length := len(opts); length%2 != 0 {
			return nil, fmt.Errorf("invalid amount of arguments [ghostty.Uid]: %v", length)
		}

		uid := &url.URL{
			Scheme: "ghostty",
			Host:   host,
			Path:   s,
		}
		values := uid.Query()
		for i := 0; i < len(opts); i += 2 {
			if opts[i+1] != "" {
				values.Add(opts[i], opts[i+1])
			}
		}
		uid.RawQuery = values.Encode()

		return uid, nil
	}
}
