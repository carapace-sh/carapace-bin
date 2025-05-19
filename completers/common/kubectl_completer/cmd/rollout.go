package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolloutCmd = &cobra.Command{
	Use:     "rollout SUBCOMMAND",
	Short:   "Manage the rollout of a resource",
	GroupID: "deploy",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolloutCmd).Standalone()

	rootCmd.AddCommand(rolloutCmd)
}
