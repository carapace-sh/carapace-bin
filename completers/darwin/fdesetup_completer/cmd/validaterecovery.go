package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var validaterecoveryCmd = &cobra.Command{
	Use:   "validaterecovery",
	Short: "Check if the recovery key is valid",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(validaterecoveryCmd).Standalone()
	rootCmd.AddCommand(validaterecoveryCmd)
}
