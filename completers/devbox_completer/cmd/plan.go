package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var planCmd = &cobra.Command{
	Use:   "plan",
	Short: "Preview the plan used to build your environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(planCmd).Standalone()
	rootCmd.AddCommand(planCmd)

	carapace.Gen(planCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
