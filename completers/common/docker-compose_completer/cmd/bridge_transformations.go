package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bridge_transformationsCmd = &cobra.Command{
	Use:   "transformations CMD [OPTIONS]",
	Short: "Manage transformation images",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bridge_transformationsCmd).Standalone()

	bridgeCmd.AddCommand(bridge_transformationsCmd)
}
