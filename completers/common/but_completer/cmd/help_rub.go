package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_rubCmd = &cobra.Command{
	Use:   "rub",
	Short: "Combines two entities together to perform an operation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_rubCmd).Standalone()

	helpCmd.AddCommand(help_rubCmd)
}
