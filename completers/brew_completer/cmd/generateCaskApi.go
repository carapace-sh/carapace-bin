package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var generateCaskApiCmd = &cobra.Command{
	Use:     "generate-cask-api",
	Short:   "Generate `homebrew/cask` API data files for <https://formulae.brew.sh>",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generateCaskApiCmd).Standalone()

	generateCaskApiCmd.Flags().Bool("debug", false, "Display any debugging information.")
	generateCaskApiCmd.Flags().Bool("dry-run", false, "Generate API data without writing it to files.")
	generateCaskApiCmd.Flags().Bool("help", false, "Show this message.")
	generateCaskApiCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	generateCaskApiCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	rootCmd.AddCommand(generateCaskApiCmd)
}
