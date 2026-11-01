package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rollbackCmd = &cobra.Command{
	Use:     "rollback",
	Short:   "roll back the last transaction (DANGEROUS) (DEPRECATED)",
	GroupID: groups[group_repository_maintenance].ID,
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rollbackCmd).Standalone()

	rollbackCmd.Flags().BoolP("dry-run", "n", false, "do not perform actions, just print output")
	rollbackCmd.Flags().BoolP("force", "f", false, "ignore safety measures")
	rootCmd.AddCommand(rollbackCmd)
}
