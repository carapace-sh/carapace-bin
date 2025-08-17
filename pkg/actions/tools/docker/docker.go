// package docker contains docker related actions
package docker

import (
	"fmt"
	"net/url"

	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/styles"
	"github.com/carapace-sh/carapace/pkg/uid"
)

// Uid TODO experimental
func Uid(host string, opts ...string) func(s string, uc uid.Context) (*url.URL, error) {
	return func(s string, uc uid.Context) (*url.URL, error) {
		if length := len(opts); length%2 != 0 {
			return nil, fmt.Errorf("invalid amount of arguments [docker.Uid]: %v", length)
		}

		uid := &url.URL{
			Scheme: "docker",
			Host:   host,
			Path:   s,
		}
		values := uid.Query()
		if e := uc.Getenv("DOCKER_HOST"); e != "" {
			values.Add("DOCKER_HOST", e)
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

// ActionRepositories completes repository names
//
//	alpine
//	bash
func ActionRepositories() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "images", "--format", "{{.Repository}}")(func(output []byte) carapace.Action {
			vals := strings.Split(string(output), "\n")
			return carapace.ActionValues(vals[:len(vals)-1]...)
		})
	}).Tag("repositories")
}

// ActionRepositoryTags completes repository names and tags separately
//
//	alpine:latest
//	bash:latest
func ActionRepositoryTags() carapace.Action {
	return carapace.ActionMultiParts(":", func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("docker", "image", "ls", "--format", "{{.Repository}}:{{.Tag}}", "--filter", "dangling=false")(func(output []byte) carapace.Action {
			// TODO add checks to not cause an out of bounds error
			lines := strings.Split(string(output), "\n")
			switch len(c.Parts) {
			case 0:
				reposWithSuffix := make([]string, len(lines)-1)
				for index, val := range lines[:len(lines)-1] {
					if val != " " {
						reposWithSuffix[index] = strings.SplitAfter(val, ":")[0]
					}
				}
				return carapace.ActionValues(reposWithSuffix...).Tag("repositories")
			case 1:
				tags := make([]string, 0)
				for _, val := range lines[:len(lines)-1] {
					if strings.HasPrefix(val, c.Parts[0]) {
						tag := strings.Split(val, ":")[1]
						tags = append(tags, tag)
					}
				}
				return carapace.ActionValues(tags...).Tag("tags")
			default:
				return carapace.ActionValues()
			}
		}).Style(styles.Docker.Image)
	})
}
