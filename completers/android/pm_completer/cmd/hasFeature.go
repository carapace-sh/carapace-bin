package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var hasFeatureCmd = &cobra.Command{
	Use:   "has-feature",
	Short: "Prints true when the system has the given feature",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hasFeatureCmd).Standalone()

	rootCmd.AddCommand(hasFeatureCmd)

	carapace.Gen(hasFeatureCmd).PositionalAnyCompletion(
		android.ActionFeatures(),
	)
}
