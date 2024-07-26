package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var horizonSupervisorsCmd = &cobra.Command{
	Use:   "horizon:supervisors",
	Short: "List all of the supervisors",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(horizonSupervisorsCmd).Standalone()

	rootCmd.AddCommand(horizonSupervisorsCmd)
}
