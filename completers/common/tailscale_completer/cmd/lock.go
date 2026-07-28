package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lockCmd = &cobra.Command{
	Use:   "lock",
	Short: "Manage tailnet lock",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lockCmd).Standalone()

	rootCmd.AddCommand(lockCmd)
}
