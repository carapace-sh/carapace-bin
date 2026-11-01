package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ntable_showCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"list"},
	Short:   "list neighbour table parameters and statistics",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ntable_showCmd).Standalone()

	ntableCmd.AddCommand(ntable_showCmd)
}
