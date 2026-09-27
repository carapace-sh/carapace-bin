package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_rbCmd = &cobra.Command{
	Use:   "rb",
	Short: "Rebuild a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_rbCmd).Standalone()

	helpCmd.AddCommand(help_rbCmd)
}
