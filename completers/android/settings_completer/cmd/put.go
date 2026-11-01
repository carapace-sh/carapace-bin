package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/android/settings_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var putCmd = &cobra.Command{
	Use:   "put",
	Short: "Change the contents of a key to a value",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(putCmd).Standalone()

	putCmd.Flags().String("user", "", "specify which `USER_ID` to modify")

	rootCmd.AddCommand(putCmd)

	carapace.Gen(putCmd).FlagCompletion(carapace.ActionMap{
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(putCmd).PositionalCompletion(
		action.ActionNamespaces(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionKeys(c.Args[0])
		}),
		carapace.ActionValues(),          // VALUE
		carapace.ActionValues("default"), // TAG or default
		carapace.ActionValues("default"),
	)
}
