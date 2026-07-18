package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bumpCompatibilityVersionCmd = &cobra.Command{
	Use:     "bump-compatibility-version",
	Short:   "Create a commit to increment the compatibility_version of <formula>",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bumpCompatibilityVersionCmd).Standalone()

	bumpCompatibilityVersionCmd.Flags().Bool("debug", false, "Display any debugging information.")
	bumpCompatibilityVersionCmd.Flags().Bool("dry-run", false, "Print what would be done rather than doing it.")
	bumpCompatibilityVersionCmd.Flags().Bool("help", false, "Show this message.")
	bumpCompatibilityVersionCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	bumpCompatibilityVersionCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	bumpCompatibilityVersionCmd.Flags().Bool("write-only", false, "Make the expected file modifications without taking any Git actions.")
	rootCmd.AddCommand(bumpCompatibilityVersionCmd)
}
