package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keygenCmd = &cobra.Command{
	Use:   "keygen",
	Short: "Generates a new encryption key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keygenCmd).Standalone()

	rootCmd.AddCommand(keygenCmd)
}
