package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/android/settings_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve the current value of a key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getCmd).Standalone()

	getCmd.Flags().String("user", "", "specify which `USER_ID` to query")

	rootCmd.AddCommand(getCmd)

	carapace.Gen(getCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(getCmd).PositionalCompletion(
		action.ActionNamespaces(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionKeys(c.Args[0])
		}),
	)
}
