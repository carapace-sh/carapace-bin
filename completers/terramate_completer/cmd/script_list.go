package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var script_listCmd = &cobra.Command{
	Use:   "list",
	Short: "Show a list of all scripts in the current directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(script_listCmd).Standalone()

	scriptCmd.AddCommand(script_listCmd)
}
