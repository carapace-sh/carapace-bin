package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var pennantFeatureCmd = &cobra.Command{
	Use:   "pennant:feature [-f|--force] [--] <name>",
	Short: "Create a new feature class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pennantFeatureCmd).Standalone()

	pennantFeatureCmd.Flags().Bool("force", false, "Create the class even if the feature already exists")
	rootCmd.AddCommand(pennantFeatureCmd)
}
