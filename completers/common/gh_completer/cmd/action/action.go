package action

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action/config"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action/ghrepo"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action/git"
	"github.com/spf13/cobra"
)

func ghExecutable() string {
	return "gh"
}

func ActionConfigHosts() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if config, err := config.ParseDefaultConfig(); err != nil {
			return carapace.ActionMessage("failed to parse DefaultConfig: " + err.Error())
		} else {
			if hosts, err := config.Hosts(); err != nil {
				return carapace.ActionMessage("failed ot load hosts: " + err.Error())
			} else {
				return carapace.ActionValues(hosts...)
			}
		}
	}).Tag("config hosts")
}

func ApiV3Action(cmd *cobra.Command, query string, v interface{}, transform func() carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if repo, err := repoOverride(cmd, c); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			args := []string{
				"api",
				"--preview", "mercy",
				query,
			}
			if repo.RepoHost() != "" {
				args = append(args, "--hostname", repo.RepoHost())
			}
			return carapace.ActionExecCommand("gh", args...)(func(output []byte) carapace.Action {
				if err := json.Unmarshal(output, &v); err != nil {
					return carapace.ActionMessage("failed to unmarshall response: " + err.Error())
				}
				return transform()
			})
		}
	})
}

func GraphQlAction(cmd *cobra.Command, query string, v interface{}, transform func() carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
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

		if repo, err := repoOverride(cmd, c); err != nil {
			return carapace.ActionMessage(err.Error())
		} else {
			args := []string{
				"api", "graphql",
				"--header", "Accept: application/vnd.github.merge-info-preview+json",
				"-F", "owner=" + repo.RepoOwner(),
				"-F", "repo=" + repo.RepoName(),
				"-f", fmt.Sprintf("query=query%v {%v}", queryParams, query),
			}
			if repo.RepoHost() != "" {
				args = append(args, "--hostname", repo.RepoHost())
			}
			return carapace.ActionExecCommand(ghExecutable(), args...)(func(output []byte) carapace.Action {
				if err := json.Unmarshal(output, &v); err != nil {
					return carapace.ActionMessage("failed to unmarshall response: " + err.Error())
				}
				return transform()
			})
		}
	})
}

func repoOverride(cmd *cobra.Command, c carapace.Context) (ghrepo.Interface, error) {
	repoOverride := ""
	if flag := cmd.Flag("repo"); flag != nil {
		repoOverride = flag.Value.String()
	}
	if repoFromEnv := c.Getenv("GH_REPO"); repoOverride == "" && repoFromEnv != "" {
		repoOverride = repoFromEnv
	}
	if repoOverride != "" {
		splitted := strings.Split(repoOverride, "/")
		switch len(splitted) {
		case 1:
			return ghrepo.New(splitted[0], ""), nil // assume owner
		case 2:
			if strings.Contains(splitted[0], ".") {
				return ghrepo.NewWithHost(splitted[1], "", splitted[0]), nil
			} else {
				return ghrepo.New(splitted[0], splitted[1]), nil
			}
		default:
			return ghrepo.NewWithHost(splitted[1], splitted[2], splitted[0]), nil
		}
	}

	if remotes, err := git.Remotes(); err == nil {
		for _, remote := range remotes {
			if remote.Resolved == "base" {
				return ghrepo.FromURL(remote.FetchURL) // TODO ssh translator
			}
		}
	}
	return nil, errors.New("no default repository defined")
}

// https://yourbasic.org/golang/formatting-byte-size-to-human-readable-format/
func byteCountSI(b int) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB",
		float64(b)/float64(div), "kMGTPE"[exp])
}

func defaultUser() (user string, err error) {
	var cfg config.Config
	if cfg, err = config.ParseDefaultConfig(); err == nil {
		var hosts []string
		if hosts, err = cfg.Hosts(); err == nil {
			if len(hosts) < 1 {
				err = errors.New("could not retrieve default user")
			} else {
				user, err = cfg.Get(hosts[0], "user")
			}
		}
	}
	return
}
