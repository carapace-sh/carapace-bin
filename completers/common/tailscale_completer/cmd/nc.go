package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/tailscale"
	"github.com/spf13/cobra"
)

var ncCmd = &cobra.Command{
	Use:   "nc",
	Short: "Connect to a port on a host, connected to stdin/stdout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ncCmd).Standalone()

	rootCmd.AddCommand(ncCmd)

	carapace.Gen(ncCmd).PositionalCompletion(
		tailscale.ActionHosts(),
	)
}
