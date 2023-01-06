package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var stackCmd = &cobra.Command{
	Use:     "stack [OPTIONS]",
	Short:   "Manage Swarm stacks",
	GroupID: "management",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stackCmd).Standalone()
	rootCmd.AddCommand(stackCmd)
}
