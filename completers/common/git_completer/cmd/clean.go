package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:     "clean",
	Short:   "Remove untracked files from the working tree",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_main].ID,
}

func init() {
	carapace.Gen(cleanCmd).Standalone()

	cleanCmd.Flags().BoolS("X", "X", false, "remove only ignored files")
	cleanCmd.Flags().BoolS("d", "d", false, "remove whole directories")
	cleanCmd.Flags().BoolP("dry-run", "n", false, "dry run")
	cleanCmd.Flags().StringP("exclude", "e", "", "add <pattern> to ignore rules")
	cleanCmd.Flags().BoolP("force", "f", false, "force")
	cleanCmd.Flags().BoolP("interactive", "i", false, "interactive cleaning")
	cleanCmd.Flags().BoolP("quiet", "q", false, "do not print names of files removed")
	cleanCmd.Flags().BoolS("x", "x", false, "remove ignored files, too")
	rootCmd.AddCommand(cleanCmd)

	carapace.Gen(cleanCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			opts := git.ChangeOpts{
				Ignored:  cleanCmd.Flag("x").Changed || cleanCmd.Flag("X").Changed,
				Unstaged: !cleanCmd.Flag("X").Changed,
			}
			return git.ActionChanges(opts).FilterArgs()
		}),
	)

	carapace.Gen(cleanCmd).DashAnyCompletion(
		carapace.ActionPositional(cleanCmd),
	)
}
