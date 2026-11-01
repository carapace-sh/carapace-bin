package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_binCmd = &cobra.Command{
	Use:   "bin",
	Short: "Print the directory where pnpm will install executables",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_binCmd).Standalone()

	helpCmd.AddCommand(help_binCmd)
}
