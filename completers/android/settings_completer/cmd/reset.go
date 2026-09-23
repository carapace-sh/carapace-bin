package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var resetCmd = &cobra.Command{
	Use:   "reset",
	Short: "Reset the global/secure table for a package with a mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resetCmd).Standalone()

	resetCmd.Flags().String("user", "", "specify which `USER_ID` to modify")

	rootCmd.AddCommand(resetCmd)

	carapace.Gen(resetCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(resetCmd).PositionalCompletion(
		carapace.ActionValues("secure", "global"),
		carapace.Batch(
			carapace.ActionValuesDescribed(
				"untrusted_defaults", "RESET_MODE_UNTRUSTED_DEFAULTS",
				"untrusted_clear", "RESET_MODE_UNTRUSTED_CHANGES",
				"trusted_defaults", "RESET_MODE_TRUSTED_DEFAULTS",
			),
			android.ActionPackages(),
		).ToA(),
	)
}
