package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var crons_addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a cron job to vercel.json",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(crons_addCmd).Standalone()

	crons_addCmd.Flags().String("path", "", "Path for the cron job")
	crons_addCmd.Flags().String("schedule", "", "Schedule for the cron job")

	cronsCmd.AddCommand(crons_addCmd)
}
