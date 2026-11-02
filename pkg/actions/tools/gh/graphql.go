package gh

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/cache"
	"github.com/carapace-sh/carapace/pkg/cache/key"
)

func graphQlAction(opts RepoOpts, query string, v any, transform func() carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		needsRepo := strings.Contains(query, "$owner") || strings.Contains(query, "$repo")
		if needsRepo {
			return resolveRepo(c, opts, func(resolved RepoOpts) carapace.Action {
				return executeGraphQl(resolved, query, v, transform)
			})
		}
		return executeGraphQl(opts, query, v, transform)
	})
}

func executeGraphQl(opts RepoOpts, query string, v any, transform func() carapace.Action) carapace.Action {
	return actionHostConfig(func(config hostConfig) carapace.Action {
		params := make([]string, 0)
		if strings.Contains(query, "$owner") {
			params = append(params, "$owner: String!")
		}
		if strings.Contains(query, "$repo") {
			params = append(params, "$repo: String!")
		}
		queryParams := strings.Join(params, ",")
		if queryParams != "" {
			queryParams = "(" + queryParams + ")"
		}

		if opts.Owner == "@me" {
			conf, ok := config[opts.Host]
			if !ok {
				return carapace.ActionMessage("unknown host")
			}
			opts.Owner = conf.User
		}

		args := []string{"api", "graphql",
			"--header", "Accept: application/vnd.github.merge-info-preview+json",
		}
		if strings.Contains(query, "$owner") {
			args = append(args, "-F", "owner="+opts.Owner)
		}
		if strings.Contains(query, "$repo") {
			args = append(args, "-F", "repo="+opts.Name)
		}
		args = append(args, "-f", fmt.Sprintf("query=query%v {%v}", queryParams, query))

		if opts.Host != "" {
			args = append(args, "--hostname", opts.Host)
		}
		return carapace.ActionExecCommand("gh", args...)(func(output []byte) carapace.Action {
			if err := json.Unmarshal(output, &v); err != nil {
				return carapace.ActionMessage("failed to unmarshall response: " + err.Error())
			}
			return transform()
		})
	})
}

// resolveRepo resolves the repository owner and name when they are empty,
// using `gh repo view` which performs the same implicit repo detection as
// the gh CLI (git remotes, gh-resolved config, GH_REPO env var).
// The hostname is not covered by that: `gh api` ignores the repository
// context (it only honors --hostname or GH_HOST), so the host part of
// GH_REPO is parsed here to keep enterprise hosts working.
func resolveRepo(c carapace.Context, opts RepoOpts, f func(RepoOpts) carapace.Action) carapace.Action {
	if opts.Owner != "" && opts.Name != "" {
		return f(opts)
	}

	if opts.Host == "" {
		if envRepo := c.Getenv("GH_REPO"); envRepo != "" {
			splitted := strings.Split(envRepo, "/")
			if len(splitted) > 2 && strings.Contains(splitted[0], ".") {
				opts.Host = splitted[0] // enterprise host given with GH_REPO=HOST/OWNER/REPO
			}
		}
	}

	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		content, err := cache.Cache(24*time.Hour,
			key.String("gh", "repoView", opts.Host, opts.Owner, opts.Name, c.Getenv("GH_REPO"), c.Dir),
		)(func() ([]byte, error) {
			var stdout, stderr bytes.Buffer
			cmd := c.Command("gh", "repo", "view", "--json", "owner,name")
			cmd.Stdout = &stdout
			cmd.Stderr = &stderr
			if err := cmd.Run(); err != nil {
				if firstLine := strings.SplitN(stderr.String(), "\n", 2)[0]; strings.TrimSpace(firstLine) != "" {
					err = errors.New(firstLine)
				}
				return nil, err
			}
			return stdout.Bytes(), nil
		})
		if err != nil {
			return carapace.ActionMessage("failed to resolve repository: " + err.Error())
		}

		var identity struct {
			Owner struct {
				Login string `json:"login"`
			} `json:"owner"`
			Name string `json:"name"`
		}
		if err := json.Unmarshal(content, &identity); err != nil {
			return carapace.ActionMessage("failed to resolve repository: " + err.Error())
		}
		if opts.Owner == "" {
			opts.Owner = identity.Owner.Login
		}
		if opts.Name == "" {
			opts.Name = identity.Name
		}
		return f(opts)
	})
}
