package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Print the path to the .apk of the given package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pathCmd).Standalone()

	pathCmd.Flags().String("user", "", "specify which `USER_ID` to query")

	rootCmd.AddCommand(pathCmd)

	carapace.Gen(pathCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(pathCmd).PositionalAnyCompletion(
		android.ActionPackages(),
	)
}
