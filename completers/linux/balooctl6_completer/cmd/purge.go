package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "Remove the index database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(purgeCmd).Standalone()

	rootCmd.AddCommand(purgeCmd)
}
