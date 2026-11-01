package gh

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/carapace-sh/carapace"
)

func graphQlAction(opts RepoOpts, query string, v any, transform func() carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		needsRepo := strings.Contains(query, "$owner") || strings.Contains(query, "$repo")
		if needsRepo {
			return resolveRepo(opts, func(resolved RepoOpts) carapace.Action {
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
func resolveRepo(opts RepoOpts, f func(RepoOpts) carapace.Action) carapace.Action {
	if opts.Owner != "" && opts.Name != "" {
		return f(opts)
	}
	return carapace.ActionExecCommand("gh", "repo", "view", "--json", "owner,name")(func(output []byte) carapace.Action {
		var identity struct {
			Owner struct {
				Login string `json:"login"`
			} `json:"owner"`
			Name string `json:"name"`
		}
		if err := json.Unmarshal(output, &identity); err != nil {
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
