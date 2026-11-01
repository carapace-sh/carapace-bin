package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var killCmd = &cobra.Command{
	Use:   "kill",
	Short: "Kill all background processes associated with the given application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(killCmd).Standalone()

	killCmd.Flags().String("user", "", "specify which `USER_ID` to affect")

	rootCmd.AddCommand(killCmd)

	carapace.Gen(killCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("all", "current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(killCmd).PositionalAnyCompletion(
		android.ActionPackages(),
	)
}
