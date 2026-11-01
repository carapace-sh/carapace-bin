package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var skillCmd = &cobra.Command{
	Use:   "skill",
	Short: "Manage the agent skill",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(skillCmd).Standalone()

	rootCmd.AddCommand(skillCmd)
}
