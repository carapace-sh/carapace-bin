package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var skillStatusCmd = &cobra.Command{
	Use:   "status",
	Short: "Report what is installed in each agent",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skillStatusCmd).Standalone()

	skillCmd.AddCommand(skillStatusCmd)
}
