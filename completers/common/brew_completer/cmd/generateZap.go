package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/brew"
	"github.com/spf13/cobra"
)

var generateZapCmd = &cobra.Command{
	Use:     "generate-zap",
	Short:   "Generate a `zap` stanza for a cask by scanning the system for associated files and directories",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generateZapCmd).Standalone()

	generateZapCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	generateZapCmd.Flags().BoolP("help", "h", false, "Show this message.")
	generateZapCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	generateZapCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(generateZapCmd)

	carapace.Gen(generateZapCmd).PositionalAnyCompletion(
		brew.ActionAllCasks(),
	)
}
