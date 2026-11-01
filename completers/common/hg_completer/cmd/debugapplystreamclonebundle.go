package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugapplystreamclonebundleCmd = &cobra.Command{
	Use:    "debugapplystreamclonebundle",
	Short:  "FILE",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugapplystreamclonebundleCmd).Standalone()

	rootCmd.AddCommand(debugapplystreamclonebundleCmd)
}
