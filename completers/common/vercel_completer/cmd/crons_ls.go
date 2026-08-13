package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var crons_lsCmd = &cobra.Command{
	Use:     "ls",
	Aliases: []string{"list"},
	Short:   "List all cron jobs for a project",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(crons_lsCmd).Standalone()

	crons_lsCmd.Flags().String("format", "", "Output format")
	crons_lsCmd.Flags().Bool("json", false, "Output as JSON")
	crons_lsCmd.Flags().String("project", "", "Project name or ID")

	cronsCmd.AddCommand(crons_lsCmd)
}
