package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var git_colocationCmd = &cobra.Command{
	Use:   "colocation",
	Short: "Manage Jujutsu repository colocation with Git",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_colocationCmd).Standalone()

	git_colocationCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	gitCmd.AddCommand(git_colocationCmd)
}
