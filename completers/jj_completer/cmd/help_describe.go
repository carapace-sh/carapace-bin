package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var help_describeCmd = &cobra.Command{
	Use:   "describe",
	Short: "Update the change description or other metadata",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_describeCmd).Standalone()

	helpCmd.AddCommand(help_describeCmd)
}
