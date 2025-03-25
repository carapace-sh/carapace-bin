package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_testCmd = &cobra.Command{
	Use:   "test",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_testCmd).Standalone()

	helpCmd.AddCommand(help_testCmd)
}
