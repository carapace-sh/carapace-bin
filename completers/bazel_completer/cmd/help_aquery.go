package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_aqueryCmd = &cobra.Command{
	Use:   "aquery",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_aqueryCmd).Standalone()

	helpCmd.AddCommand(help_aqueryCmd)
}
