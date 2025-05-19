package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swarm_joinCmd = &cobra.Command{
	Use:   "join [OPTIONS] HOST:PORT",
	Short: "Join a swarm as a node and/or manager",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swarm_joinCmd).Standalone()

	swarm_joinCmd.Flags().String("advertise-addr", "", "Advertised address (format: \"<ip|interface>[:port]\")")
	swarm_joinCmd.Flags().String("availability", "active", "Availability of the node (\"active\", \"pause\", \"drain\")")
	swarm_joinCmd.Flags().String("data-path-addr", "", "Address or interface to use for data path traffic (format: \"<ip|interface>\")")
	swarm_joinCmd.Flags().String("listen-addr", "", "Listen address (format: \"<ip|interface>[:port]\")")
	swarm_joinCmd.Flags().String("token", "", "Token for entry into the swarm")
	swarmCmd.AddCommand(swarm_joinCmd)
}
