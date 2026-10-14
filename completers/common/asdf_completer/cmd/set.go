package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/asdf_completer/cmd/action"
	"github.com/spf13/cobra"
)

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the package version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setCmd).Standalone()

	setCmd.Flags().BoolP("home", "u", false, "The version should be set in the current users home directory")
	setCmd.Flags().BoolP("parent", "p", false, "The version should be set in the closest existing .tool-versions file in a parent directory")
	rootCmd.AddCommand(setCmd)

	carapace.Gen(setCmd).PositionalCompletion(
		action.ActionPlugins(),
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.Batch(
				action.ActionInstalledVersions(c.Args[0]),
				carapace.ActionValues("latest", "latest:"),
			).ToA()
		}),
	)
}
