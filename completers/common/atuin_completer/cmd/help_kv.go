package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_kvCmd = &cobra.Command{
	Use:   "kv",
	Short: "Get or set small key-value pairs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_kvCmd).Standalone()

	helpCmd.AddCommand(help_kvCmd)
}
