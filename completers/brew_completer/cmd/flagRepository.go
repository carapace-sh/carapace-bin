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

	flagRepositoryCmd.Flags().Bool("debug", false, "Display any debugging information.")
	flagRepositoryCmd.Flags().Bool("help", false, "Show this message.")
	flagRepositoryCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	flagRepositoryCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
}
