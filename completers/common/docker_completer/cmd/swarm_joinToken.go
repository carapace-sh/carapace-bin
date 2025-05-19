package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var swarm_joinTokenCmd = &cobra.Command{
	Use:   "join-token [OPTIONS] (worker|manager)",
	Short: "Manage join tokens",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(swarm_joinTokenCmd).Standalone()

	swarm_joinTokenCmd.Flags().BoolP("quiet", "q", false, "Only display token")
	swarm_joinTokenCmd.Flags().Bool("rotate", false, "Rotate join token")
	swarmCmd.AddCommand(swarm_joinTokenCmd)

	carapace.Gen(swarm_joinTokenCmd).PositionalCompletion(
		carapace.ActionValues("worker", "manager"),
	)
}
