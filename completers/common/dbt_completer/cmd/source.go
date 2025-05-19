package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sourceCmd = &cobra.Command{
	Use:   "source",
	Short: "Manage your project's sources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sourceCmd).Standalone()

	rootCmd.AddCommand(sourceCmd)
}
