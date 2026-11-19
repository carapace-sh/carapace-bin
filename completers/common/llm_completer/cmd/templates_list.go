package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var templates_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available prompt templates",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(templates_listCmd).Standalone()

	templatesCmd.AddCommand(templates_listCmd)
}
