package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_show_globalsCmd = &cobra.Command{
	Use:   "globals",
	Short: "List globals for all stacks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_show_globalsCmd).Standalone()

	debug_showCmd.AddCommand(debug_show_globalsCmd)
}
