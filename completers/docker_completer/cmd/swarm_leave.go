package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swarm_leaveCmd = &cobra.Command{
	Use:   "leave [OPTIONS]",
	Short: "Leave the swarm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swarm_leaveCmd).Standalone()

	swarm_leaveCmd.Flags().BoolP("force", "f", false, "Force this node to leave the swarm, ignoring warnings")
	swarmCmd.AddCommand(swarm_leaveCmd)
}
