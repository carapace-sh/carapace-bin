package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var whichCmd = &cobra.Command{
	Use:   "which",
	Short: "Shows the path to the tool's executable",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(whichCmd).Standalone()

	whichCmd.Flags().Bool("plugin", false, "Show plugin path instead of tool executable")
	whichCmd.Flags().StringP("tool", "t", "", "Use a specific tool@version")
	whichCmd.Flags().Bool("version", false, "Show tool version")
	rootCmd.AddCommand(whichCmd)

	carapace.Gen(whichCmd).FlagCompletion(carapace.ActionMap{
		"tool": action.ActionInstalledToolVersions(),
	})

	carapace.Gen(whichCmd).PositionalCompletion(
		action.ActionTools(),
	)
}
