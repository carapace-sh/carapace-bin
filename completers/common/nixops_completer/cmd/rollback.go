package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var RollbackCmd = &cobra.Command{
	Use:   "rollback",
	Short: "Rollback",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(RollbackCmd).Standalone()
	rootCmd.AddCommand(RollbackCmd)
}
