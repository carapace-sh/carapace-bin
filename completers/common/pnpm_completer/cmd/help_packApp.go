package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_packAppCmd = &cobra.Command{
	Use:   "pack-app",
	Short: "Pack a `CommonJS` entry file into a standalone executable for one or more target platforms",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_packAppCmd).Standalone()

	helpCmd.AddCommand(help_packAppCmd)
}
