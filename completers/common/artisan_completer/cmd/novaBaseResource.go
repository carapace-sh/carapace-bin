package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var novaBaseResourceCmd = &cobra.Command{
	Use:   "nova:base-resource <name>",
	Short: "Create a new base resource class",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(novaBaseResourceCmd).Standalone()

	rootCmd.AddCommand(novaBaseResourceCmd)
}
