package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_prefixCmd = &cobra.Command{
	Use:   "prefix",
	Short: "Print the current package prefix",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_prefixCmd).Standalone()

	helpCmd.AddCommand(help_prefixCmd)
}
