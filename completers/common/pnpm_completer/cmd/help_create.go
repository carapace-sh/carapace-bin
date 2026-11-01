package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_createCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates a project from a `create-*` starter kit",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_createCmd).Standalone()

	helpCmd.AddCommand(help_createCmd)
}
