package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var skillStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Report whether Rex's skill and hooks are installed",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skillStatusCmd).Standalone()

	skillStatusCmd.Flags().Bool("json", false, "print the report as JSON")
	skillStatusCmd.Flags().Bool("project", false, "report on the current repository instead of your account")

	skillCmd.AddCommand(skillStatusCmd)
}
