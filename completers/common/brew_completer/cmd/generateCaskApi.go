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

	generateCaskApiCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	generateCaskApiCmd.Flags().BoolP("dry-run", "n", false, "Generate API data without writing it to files.")
	generateCaskApiCmd.Flags().BoolP("help", "h", false, "Show this message.")
	generateCaskApiCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	generateCaskApiCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	rootCmd.AddCommand(generateCaskApiCmd)
}
