package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_resolveCmd = &cobra.Command{
	Use:   "resolve",
	Short: "Resolve revision specs to object ids and types",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_resolveCmd).Standalone()

	helpCmd.AddCommand(help_resolveCmd)
}
