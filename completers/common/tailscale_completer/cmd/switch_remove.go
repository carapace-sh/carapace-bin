package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var switch_removeCmd = &cobra.Command{
	Use:   "remove",
	Short: "Remove a Tailscale account",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(switch_removeCmd).Standalone()

	switchCmd.AddCommand(switch_removeCmd)
}
