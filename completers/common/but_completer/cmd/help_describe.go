package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Edit the commit message of the specified commit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_describeCmd).Standalone()

	helpCmd.AddCommand(help_describeCmd)
}
