package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var novaResourceToolCmd = &cobra.Command{
	Use:   "nova:resource-tool <name>",
	Short: "Create a new resource tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(novaResourceToolCmd).Standalone()

	rootCmd.AddCommand(novaResourceToolCmd)
}
