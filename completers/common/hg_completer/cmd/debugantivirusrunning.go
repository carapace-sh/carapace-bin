package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugantivirusrunningCmd = &cobra.Command{
	Use:    "debugantivirusrunning",
	Short:  "attempt to trigger an antivirus scanner to see if one is active",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugantivirusrunningCmd).Standalone()

	rootCmd.AddCommand(debugantivirusrunningCmd)
}
