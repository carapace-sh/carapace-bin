package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var novaToolCmd = &cobra.Command{
	Use:   "nova:tool <name>",
	Short: "Create a new tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(novaToolCmd).Standalone()

	rootCmd.AddCommand(novaToolCmd)
}
