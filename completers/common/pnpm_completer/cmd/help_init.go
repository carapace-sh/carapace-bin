package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize a package.json",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_initCmd).Standalone()

	helpCmd.AddCommand(help_initCmd)
}
