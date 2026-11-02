package action

import (
	"bytes"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh/config"
	"github.com/carapace-sh/carapace/pkg/execlog"
	"github.com/spf13/cobra"
)

// repoOpts parses the `--repo` flag (`[HOST/]OWNER/REPO`) into gh.RepoOpts.
// Missing fields are resolved by the corresponding public action,
// `GH_REPO` and git-remotes detection happen inside `gh repo view`.
func repoOpts(cmd *cobra.Command) gh.RepoOpts {
	if flag := cmd.Flag("repo"); flag != nil {
		value := flag.Value.String()
		if value != "" {
			splitted := strings.Split(value, "/")
			switch len(splitted) {
			case 1:
				return gh.RepoOpts{Owner: splitted[0]} // assume owner
			case 2:
				if strings.Contains(splitted[0], ".") {
					return gh.RepoOpts{Host: splitted[0], Name: splitted[1]}
				}
				return gh.RepoOpts{Owner: splitted[0], Name: splitted[1]}
			default:
				return gh.RepoOpts{Host: splitted[0], Owner: splitted[1], Name: splitted[2]}
			}
		}
	}
	return gh.RepoOpts{}
}

func ActionConfigHosts() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if config, err := config.ParseDefaultConfig(); err != nil {
			return carapace.ActionMessage("failed to parse DefaultConfig: " + err.Error())
		} else {
			if hosts, err := config.Hosts(); err != nil {
				return carapace.ActionMessage("failed to load hosts: " + err.Error())
			} else {
				return carapace.ActionValues(hosts...)
			}
		}
	}).Tag("config hosts")
}

type Commit struct {
	Sha   string
	Title string
}

// LastCommit returns the sha and title of the HEAD commit of git.
func LastCommit() (*Commit, error) {
	cmd := execlog.Command("git", "-c", "log.ShowSignature=false", "show", "-s", "--pretty=format:%H,%s", "HEAD")
	var stdout bytes.Buffer
	cmd.Stdout = &stdout
	if err := cmd.Run(); err != nil {
		return nil, err
	}

	output := stdout.Bytes()
	idx := bytes.IndexByte(output, ',')
	return &Commit{
		Sha:   string(output[0:idx]),
		Title: strings.TrimSpace(string(output[idx+1:])),
	}, nil
}
