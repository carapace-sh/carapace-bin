package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var toolAlias_lsCmd = &cobra.Command{
	Use:     "ls",
	Short:   "List tool version aliases",
	Aliases: []string{"list"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(toolAlias_lsCmd).Standalone()

	toolAlias_lsCmd.Flags().BoolP("help", "h", false, "Print help")
	toolAlias_lsCmd.Flags().Bool("no-header", false, "Do not print the table header")
	toolAliasCmd.AddCommand(toolAlias_lsCmd)

	carapace.Gen(toolAlias_lsCmd).PositionalCompletion(
		mise.ActionTools(),
	)
}
