package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_self_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Download and install updates to rustup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_self_updateCmd).Standalone()

	help_selfCmd.AddCommand(help_self_updateCmd)
}
