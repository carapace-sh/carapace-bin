package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var flagRepositoryCmd = &cobra.Command{
	Use:   "repository",
	Short: "Display where Homebrew's Git repository is located",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(flagRepositoryCmd).Standalone()

	flagRepositoryCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	flagRepositoryCmd.Flags().BoolP("help", "h", false, "Show this message.")
	flagRepositoryCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	flagRepositoryCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
}
