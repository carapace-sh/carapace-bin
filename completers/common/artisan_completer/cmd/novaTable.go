package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var novaTableCmd = &cobra.Command{
	Use:   "nova:table <name>",
	Short: "Create a new metric (table) class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(novaTableCmd).Standalone()

	rootCmd.AddCommand(novaTableCmd)
}
