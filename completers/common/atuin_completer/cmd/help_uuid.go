package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_uuidCmd = &cobra.Command{
	Use:   "uuid",
	Short: "Generate a UUID",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_uuidCmd).Standalone()

	helpCmd.AddCommand(help_uuidCmd)
}
