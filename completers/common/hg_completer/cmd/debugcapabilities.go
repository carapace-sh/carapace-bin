package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugcapabilitiesCmd = &cobra.Command{
	Use:    "debugcapabilities",
	Short:  "PATH",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugcapabilitiesCmd).Standalone()

	debugcapabilitiesCmd.Flags().String("bundle2-cap", "", "only display the matching capabilities")
	rootCmd.AddCommand(debugcapabilitiesCmd)
}
