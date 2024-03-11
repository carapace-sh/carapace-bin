package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var tunnelCmd = &cobra.Command{
	Use:     "tunnel",
	Short:   "Connect to LoadBalancer services",
	GroupID: "networking",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tunnelCmd).Standalone()
	tunnelCmd.Flags().BoolP("cleanup", "c", true, "call with cleanup=true to remove old tunnels")
	rootCmd.AddCommand(tunnelCmd)
}
