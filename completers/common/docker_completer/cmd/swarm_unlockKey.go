package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swarm_unlockKeyCmd = &cobra.Command{
	Use:   "unlock-key [OPTIONS]",
	Short: "Manage the unlock key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swarm_unlockKeyCmd).Standalone()

	swarm_unlockKeyCmd.Flags().BoolP("quiet", "q", false, "Only display token")
	swarm_unlockKeyCmd.Flags().Bool("rotate", false, "Rotate unlock key")
	swarmCmd.AddCommand(swarm_unlockKeyCmd)
}
