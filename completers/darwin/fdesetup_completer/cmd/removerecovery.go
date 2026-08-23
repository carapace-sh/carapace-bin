package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var removerecoveryCmd = &cobra.Command{
	Use:   "removerecovery",
	Short: "Remove recovery key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(removerecoveryCmd).Standalone()

	removerecoveryCmd.Flags().Bool("verbose", false, "Enable verbose mode")
}
