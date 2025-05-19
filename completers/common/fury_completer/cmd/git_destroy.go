package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var git_destroyCmd = &cobra.Command{
	Use:     "destroy REPO",
	Short:   "Remove Git repository",
	Aliases: []string{"reset"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_destroyCmd).Standalone()

	git_destroyCmd.Flags().Bool("reset-only", false, "Reset repo without destroying")
	gitCmd.AddCommand(git_destroyCmd)
}
