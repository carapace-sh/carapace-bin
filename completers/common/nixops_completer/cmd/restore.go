package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var RestoreCmd = &cobra.Command{
	Use:   "restore",
	Short: "Restore",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(RestoreCmd).Standalone()
	rootCmd.AddCommand(RestoreCmd)
}
