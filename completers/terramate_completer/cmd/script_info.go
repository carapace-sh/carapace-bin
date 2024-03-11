package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var script_infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show detailed information about a script",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(script_infoCmd).Standalone()

	scriptCmd.AddCommand(script_infoCmd)
}
