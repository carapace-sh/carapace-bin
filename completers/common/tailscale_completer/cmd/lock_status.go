package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lock_statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Output the state of tailnet lock",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lock_statusCmd).Standalone()

	lockCmd.AddCommand(lock_statusCmd)
}
