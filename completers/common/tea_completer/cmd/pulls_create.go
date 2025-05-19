package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/tea_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tea"
	"github.com/spf13/cobra"
)

var pulls_createCmd = &cobra.Command{
	Use:     "create",
	Short:   "Create a pull-request",
	Aliases: []string{"c"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pulls_createCmd).Standalone()

	pulls_createCmd.Flags().Bool("allow-maintainer-edits", false, "Enable maintainers to push to the base branch of created pull")
	pulls_createCmd.Flags().StringP("assignees", "a", "", "Comma-separated list of usernames to assign")
	pulls_createCmd.Flags().StringP("base", "b", "", "Branch name of the PR target (default is repos default branch)")
	pulls_createCmd.Flags().StringP("deadline", "D", "", "Deadline timestamp to assign")
	pulls_createCmd.Flags().StringP("description", "d", "", "")
	pulls_createCmd.Flags().String("head", "", "Branch name of the PR source (default is current one). To specify a different head repo, use <user>:<branch>")
	pulls_createCmd.Flags().StringP("labels", "L", "", "Comma-separated list of labels to assign")
	pulls_createCmd.Flags().StringP("login", "l", "", "Use a different Gitea Login. Optional")
	pulls_createCmd.Flags().StringP("milestone", "m", "", "Milestone to assign")
	pulls_createCmd.Flags().StringP("referenced-version", "v", "", "commit-hash or tag name to assign")
	pulls_createCmd.Flags().StringP("remote", "R", "", "Discover Gitea login from remote. Optional")
	pulls_createCmd.Flags().StringP("repo", "r", "", "Override local repository path or gitea repository slug to interact with. Optional")
	pulls_createCmd.Flags().StringP("title", "t", "", "")
	pullsCmd.AddCommand(pulls_createCmd)

	// TODO completion
	carapace.Gen(pulls_createCmd).FlagCompletion(carapace.ActionMap{
		"labels": action.ActionLabels(pulls_createCmd).UniqueList(","),
		"login":  tea.ActionLogins(),
		"remote": git.ActionRemotes(),
	})
}
