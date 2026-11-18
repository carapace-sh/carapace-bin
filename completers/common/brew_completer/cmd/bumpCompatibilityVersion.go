package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/brew"
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

	bumpCompatibilityVersionCmd.Flags().BoolP("debug", "d", false, "Display any debugging information.")
	bumpCompatibilityVersionCmd.Flags().BoolP("dry-run", "n", false, "Print what would be done rather than doing it.")
	bumpCompatibilityVersionCmd.Flags().BoolP("help", "h", false, "Show this message.")
	bumpCompatibilityVersionCmd.Flags().BoolP("quiet", "q", false, "Make some output more quiet.")
	bumpCompatibilityVersionCmd.Flags().BoolP("verbose", "v", false, "Make some output more verbose.")
	bumpCompatibilityVersionCmd.Flags().Bool("write-only", false, "Make the expected file modifications without taking any Git actions.")
	rootCmd.AddCommand(bumpCompatibilityVersionCmd)

	carapace.Gen(bumpCompatibilityVersionCmd).PositionalAnyCompletion(
		brew.ActionAllFormulae(),
	)
}
