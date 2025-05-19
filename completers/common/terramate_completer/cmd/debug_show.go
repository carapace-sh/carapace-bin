package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_showCmd = &cobra.Command{
	Use:   "show",
	Short: "Show information available in the project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_showCmd).Standalone()

	debugCmd.AddCommand(debug_showCmd)
}
