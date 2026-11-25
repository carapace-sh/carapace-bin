package action

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func cachedRepos(cmd *cobra.Command, f func(repos []string) carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"cache", "ls", "--quiet"}
		if flag := cmd.Flag("cache-dir"); flag != nil && flag.Changed {
			args = append(args, "--cache-dir", flag.Value.String())
		}
		return carapace.ActionExecCommand("hf", args...)(func(output []byte) carapace.Action {
			return f(strings.Fields(string(output)))
		})
	})
}

// ActionCachedRepos completes cached repos as `<type>/<name>` (e.g. model/gpt2)
func ActionCachedRepos(cmd *cobra.Command) carapace.Action {
	return cachedRepos(cmd, func(repos []string) carapace.Action {
		return carapace.ActionValues(repos...)
	}).Tag("cached repos")
}

// ActionCachedRepoNames completes the names of cached repos matching `--repo-type`/`--type` (default model)
func ActionCachedRepoNames(cmd *cobra.Command) carapace.Action {
	return cachedRepos(cmd, func(repos []string) carapace.Action {
		repoType := "model"
		for _, name := range []string{"repo-type", "type"} {
			if flag := cmd.Flag(name); flag != nil && flag.Changed {
				repoType = flag.Value.String()
			}
		}

		prefix := repoType + "/"
		names := make([]string, 0, len(repos))
		for _, repo := range repos {
			if strings.HasPrefix(repo, prefix) {
				names = append(names, strings.TrimPrefix(repo, prefix))
			}
		}
		return carapace.ActionValues(names...)
	}).Tag("cached repos")
}
