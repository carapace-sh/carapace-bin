package action

import (
	"os"
	"regexp"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/util"
	"github.com/spf13/cobra"
)

func ActionRepoOverride(cmd *cobra.Command) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		// TODO support github enterprise (different host)
		a := carapace.ActionValues().Invoke(c)
		if wd, err := os.Getwd(); err == nil {
			if _, err := util.FindReverse(wd, ".git"); err == nil {
				a = a.Merge(actionRemoteRepositories().Invoke(c))
			}
		}
		if c.CallbackValue != "" {
			a = a.Merge(ActionOwnerRepositories(cmd).Invoke(c))
		}
		return a.ToA()
	})
}

func actionRemoteRepositories() carapace.Action {
	return carapace.ActionExecCommand("git", "remote", "--verbose")(func(output []byte) carapace.Action {
		lines := strings.Split(string(output), "\n")
		r := regexp.MustCompile(`^(?P<remote>[^ ]+)\t+(?P<repo>(?P<url>[^ ]*)\.git) (?P<type>.*)$`)

		urls := make(map[string]string)
		for _, line := range lines[:len(lines)-1] {
			if r.MatchString(line) {
				matches := r.FindStringSubmatch(line)
				urls[matches[3]] = matches[1]
			}
		}

		vals := make([]string, 0)
		for url, remote := range urls {
			vals = append(vals, url, remote)
		}
		return carapace.ActionValuesDescribed(vals...)
	})
}
