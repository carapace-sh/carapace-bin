package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/windows/scoop_completer/cmd/action"
	"github.com/spf13/cobra"
)

var unholdCmd = &cobra.Command{
	Use:   "unhold",
	Short: "unhold an app to enable updates",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(unholdCmd).Standalone()
	unholdCmd.Flags().BoolP("global", "g", false, "unhold globally installed apps")
	rootCmd.AddCommand(unholdCmd)

	carapace.Gen(unholdCmd).PositionalAnyCompletion(
		action.ActionInstalledApps(),
	)
}
