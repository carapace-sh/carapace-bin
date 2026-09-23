package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var clearCmd = &cobra.Command{
	Use:   "clear",
	Short: "Deletes data associated with a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clearCmd).Standalone()

	clearCmd.Flags().Bool("cache-only", false, "only clear cache data")
	clearCmd.Flags().String("user", "", "specifies the `USER_ID` for which we need to clear data")

	rootCmd.AddCommand(clearCmd)

	carapace.Gen(clearCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(clearCmd).PositionalCompletion(
		android.ActionPackages(),
	)
}
