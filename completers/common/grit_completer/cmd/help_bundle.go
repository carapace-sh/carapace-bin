package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_bundleCmd = &cobra.Command{
	Use:   "bundle",
	Short: "Create, verify, or list git bundle files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_bundleCmd).Standalone()

	helpCmd.AddCommand(help_bundleCmd)
}
