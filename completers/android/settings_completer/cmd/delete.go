package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/android/settings_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete the entry for a key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deleteCmd).Standalone()

	deleteCmd.Flags().String("user", "", "specify which `USER_ID` to modify")

	rootCmd.AddCommand(deleteCmd)

	carapace.Gen(deleteCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(deleteCmd).PositionalCompletion(
		action.ActionNamespaces(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionKeys(c.Args[0])
		}),
	)
}
