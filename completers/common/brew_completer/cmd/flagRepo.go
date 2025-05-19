package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var flagRepoCmd = &cobra.Command{
	Use:   "repo",
	Short: "Display where Homebrew's Git repository is located",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flagRepoCmd).Standalone()

	flagRepoCmd.Flags().Bool("debug", false, "Display any debugging information.")
	flagRepoCmd.Flags().Bool("help", false, "Show this message.")
	flagRepoCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	flagRepoCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
}
