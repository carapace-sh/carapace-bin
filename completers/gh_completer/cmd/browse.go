package cmd

import (
	"path/filepath"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var browseCmd = &cobra.Command{
	Use:   "browse [<number> | <path>]",
	Short: "Open the repository in the browser",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	browseCmd.Flags().StringP("branch", "b", "", "Select another branch by passing in the branch name")
	browseCmd.Flags().BoolP("projects", "p", false, "Open repository projects")
	browseCmd.Flags().BoolP("settings", "s", false, "Open repository settings")
	browseCmd.Flags().BoolP("wiki", "w", false, "Open repository wiki")
	browseCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository using the `[HOST/]OWNER/REPO` format")
	rootCmd.AddCommand(browseCmd)

	carapace.Gen(browseCmd).FlagCompletion(carapace.ActionMap{
		"branch": action.ActionBranches(browseCmd), // TODO merge with tags
		"repo":   action.ActionOwnerRepositories(browseCmd),
	})

	carapace.Gen(browseCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if browseCmd.Flag("projects").Changed ||
				browseCmd.Flag("settings").Changed ||
				browseCmd.Flag("wiki").Changed {
				return carapace.ActionValues()
			}

			path := filepath.Dir(c.CallbackValue)
			// TODO merge with issue/pr ids (show all with prefix search using callbackvalue)
			return carapace.ActionMultiParts("/", func(c carapace.Context) carapace.Action {
				return action.ActionContents(browseCmd, path, browseCmd.Flag("branch").Value.String())
			})
		}),
	)
}
