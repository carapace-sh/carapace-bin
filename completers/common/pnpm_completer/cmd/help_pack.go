package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Create a tarball from a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_packCmd).Standalone()

	helpCmd.AddCommand(help_packCmd)
}
