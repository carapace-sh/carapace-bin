package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugcreatestreamclonebundleCmd = &cobra.Command{
	Use:    "debugcreatestreamclonebundle",
	Short:  "FILE",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugcreatestreamclonebundleCmd).Standalone()

	rootCmd.AddCommand(debugcreatestreamclonebundleCmd)
}
