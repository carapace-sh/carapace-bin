package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugUnbundleCmd = &cobra.Command{
	Use:    "debug::unbundle",
	Short:  "FILE...",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugUnbundleCmd).Standalone()

	rootCmd.AddCommand(debugUnbundleCmd)
}
