package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var base_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Updates the workspace (with all applied branches) to include the latest changes from the base branch",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(base_updateCmd).Standalone()

	base_updateCmd.Flags().BoolP("help", "h", false, "Print help")
	baseCmd.AddCommand(base_updateCmd)
}
