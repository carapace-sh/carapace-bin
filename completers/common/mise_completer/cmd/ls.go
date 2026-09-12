package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/mise_completer/cmd/action"
	"github.com/spf13/cobra"
)

var lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List installed and available tool versions",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lsCmd).Standalone()

	lsCmd.Flags().BoolP("current", "c", false, "Only show currently active tools")
	lsCmd.Flags().BoolP("installed", "i", false, "Only show installed tools")
	lsCmd.Flags().BoolP("json", "J", false, "Output in JSON format")
	lsCmd.Flags().BoolP("offline", "m", false, "Do not fetch remote versions")
	rootCmd.AddCommand(lsCmd)

	carapace.Gen(lsCmd).PositionalCompletion(
		action.ActionTools(),
	)
}
