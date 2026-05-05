package jj

import (
	"fmt"
	"net/url"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/carapace-sh/carapace/pkg/uid"
)

func actionExecJJ(arg ...string) func(f func(output []byte) carapace.Action) carapace.Action {
	return func(f func(output []byte) carapace.Action) carapace.Action {
		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			args := []string{"--color", "never"}
			if repository, ok := c.LookupEnv("JJ_REPOSITORY"); ok {
				args = append(args, "--repository", repository)
			}
			args = append(args, arg...)
			return carapace.ActionExecCommand("jj", args...)(func(output []byte) carapace.Action {
				return f(output)
			})
		})
	}
}

func actionExecJJE(arg ...string) func(f func(output []byte, err error) carapace.Action) carapace.Action {
	return func(f func(output []byte, err error) carapace.Action) carapace.Action {
		return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			args := []string{"--color", "never"}
			if repository, ok := c.LookupEnv("JJ_REPOSITORY"); ok {
				args = append(args, "--repository", repository)
			}
			args = append(args, arg...)
			return carapace.ActionExecCommandE("jj", args...)(func(output []byte, err error) carapace.Action {
				return f(output, err)
			})
		})
	}
}

func Uid(host string, opts ...string) func(s string, uc uid.Context) (*url.URL, error) {
	return func(s string, uc uid.Context) (*url.URL, error) {
		if length := len(opts); length%2 != 0 {
			return nil, fmt.Errorf("invalid amount of arguments [jj.Uid]: %v", length)
		}

		repository, ok := uc.LookupEnv("JJ_REPOSITORY")
		if !ok {
			var err error
			repository, err = traverse.Parent(".jj")(uc)
			if err != nil {
				return nil, err
			}
		}

		uid := &url.URL{
			Scheme: "jj",
			Host:   host,
			Path:   s,
		}
		values := uid.Query()
		values.Add("JJ_REPOSITORY", repository)
		for i := 0; i < len(opts); i += 2 {
			if opts[i+1] != "" { // implicitly skip empty values
				values.Add(opts[i], opts[i+1])
			}
		}
		uid.RawQuery = values.Encode()

		return uid, nil
	}
}
