package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cronsCmd = &cobra.Command{
	Use:     "crons",
	Aliases: []string{"cron"},
	Short:   "Manage cron jobs for a project",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cronsCmd).Standalone()

	rootCmd.AddCommand(cronsCmd)
}
