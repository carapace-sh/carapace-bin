package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var whereCmd = &cobra.Command{
	Use:   "where",
	Short: "Display the installation path for a tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(whereCmd).Standalone()

	rootCmd.AddCommand(whereCmd)

	carapace.Gen(whereCmd).PositionalCompletion(
		action.ActionInstalledToolVersions(),
	)
}
