package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vrfCmd = &cobra.Command{
	Use:   "vrf",
	Short: "manage virtual routing and forwarding devices",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vrfCmd).Standalone()

	rootCmd.AddCommand(vrfCmd)
}
