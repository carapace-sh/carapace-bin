package cmd

import (
	"path/filepath"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var pr_diffCmd = &cobra.Command{
	Use:     "diff [<number> | <url> | <branch>]",
	Short:   "View changes in a pull request",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pr_diffCmd).Standalone()

	pr_diffCmd.Flags().String("color", "", "Use color in diff output: {always|never|auto}")
	pr_diffCmd.Flags().StringSliceP("exclude", "e", nil, "Exclude files matching glob `patterns` from the diff")
	pr_diffCmd.Flags().Bool("name-only", false, "Display only names of changed files")
	pr_diffCmd.Flags().Bool("patch", false, "Display diff in patch format")
	pr_diffCmd.Flags().BoolP("web", "w", false, "Open the pull request diff in the browser")
	prCmd.AddCommand(pr_diffCmd)

	carapace.Gen(pr_diffCmd).FlagCompletion(carapace.ActionMap{
		"color": carapace.ActionValues("auto", "never", "always").StyleF(style.ForKeyword),
		"exclude": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			args := []string{"pr", "diff", "--name-only"}
			if f := pr_diffCmd.Flag("repo"); f.Changed { // TODO use repoverride
				args = append(args, "--repo", f.Value.String())
			}
			args = append(args, c.Args...)
			return carapace.ActionExecCommand("gh", args...)(func(output []byte) carapace.Action {
				lines := strings.Split(string(output), "\n")
				for index, line := range lines {
					lines[index] = filepath.Base(line) // full paths don't work as flag value?
				}
				return carapace.ActionValues(lines[:len(lines)-1]...).StyleF(style.ForPathExt)
			})
		}).UniqueList(","),
	})

	carapace.Gen(pr_diffCmd).PositionalCompletion(
		action.ActionPullRequests(pr_diffCmd, action.PullRequestOpts{Open: true}),
	)
}
