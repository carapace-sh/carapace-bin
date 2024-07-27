package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var horizonPauseSupervisorCmd = &cobra.Command{
	Use:   "horizon:pause-supervisor <name>",
	Short: "Pause a supervisor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(horizonPauseSupervisorCmd).Standalone()

	rootCmd.AddCommand(horizonPauseSupervisorCmd)
}
