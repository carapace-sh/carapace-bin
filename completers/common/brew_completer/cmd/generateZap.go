package cmd

import (
	"github.com/carapace-sh/carapace"
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

	generateZapCmd.Flags().Bool("debug", false, "Display any debugging information.")
	generateZapCmd.Flags().Bool("help", false, "Show this message.")
	generateZapCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	generateZapCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(generateZapCmd)
}
