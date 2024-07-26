package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var horizonContinueSupervisorCmd = &cobra.Command{
	Use:   "horizon:continue-supervisor <name>",
	Short: "Instruct the supervisor to continue processing jobs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(horizonContinueSupervisorCmd).Standalone()

	rootCmd.AddCommand(horizonContinueSupervisorCmd)
}
