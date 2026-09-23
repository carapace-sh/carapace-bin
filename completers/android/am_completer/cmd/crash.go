package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var crashCmd = &cobra.Command{
	Use:   "crash",
	Short: "Induce a VM crash in the specified package or process",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(crashCmd).Standalone()

	crashCmd.Flags().String("user", "", "specify which `USER_ID` to affect")

	rootCmd.AddCommand(crashCmd)

	carapace.Gen(crashCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(crashCmd).PositionalAnyCompletion(
		android.ActionPackages(),
	)
}
