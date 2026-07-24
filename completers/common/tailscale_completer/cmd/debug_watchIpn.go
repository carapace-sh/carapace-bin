package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_watchIpnCmd = &cobra.Command{
	Use:   "watch-ipn",
	Short: "Subscribe to IPN message bus",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_watchIpnCmd).Standalone()

	debug_watchIpnCmd.Flags().Bool("client-version", false, "include client version changes")
	debug_watchIpnCmd.Flags().Int("count", 0, "number of updates to print before exiting")
	debug_watchIpnCmd.Flags().Bool("engine-updates", false, "include engine updates")
	debug_watchIpnCmd.Flags().Bool("health-actions", false, "include health actions")
	debug_watchIpnCmd.Flags().Bool("initial", false, "include initial state")
	debug_watchIpnCmd.Flags().Bool("initial-client-version", false, "include initial client version")
	debug_watchIpnCmd.Flags().Bool("initial-drive-shares", false, "include initial drive shares")
	debug_watchIpnCmd.Flags().Bool("initial-health", false, "include initial health")
	debug_watchIpnCmd.Flags().Bool("initial-outgoing-files", false, "include initial outgoing files")
	debug_watchIpnCmd.Flags().Bool("initial-status", false, "include initial status")
	debug_watchIpnCmd.Flags().Bool("initial-suggested-exit-node", false, "include initial suggested exit node")
	debug_watchIpnCmd.Flags().Bool("peer-changes", false, "include peer changes")
	debug_watchIpnCmd.Flags().Bool("peer-patches", false, "include peer patches")
	debug_watchIpnCmd.Flags().Bool("peer-wireguard-state", false, "include peer WireGuard state")
	debugCmd.AddCommand(debug_watchIpnCmd)
}
