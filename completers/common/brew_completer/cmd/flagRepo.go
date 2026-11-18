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

	flagRepoCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	flagRepoCmd.Flags().BoolP("help", "h", false, "Show this message.")
	flagRepoCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	flagRepoCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
}
