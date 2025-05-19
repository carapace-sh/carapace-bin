package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var novaProgressCmd = &cobra.Command{
	Use:   "nova:progress <name>",
	Short: "Create a new metric (progress) class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(novaProgressCmd).Standalone()

	rootCmd.AddCommand(novaProgressCmd)
}
