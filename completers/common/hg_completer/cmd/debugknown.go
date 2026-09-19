package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugknownCmd = &cobra.Command{
	Use:    "debugknown",
	Short:  "REPO ID...",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugknownCmd).Standalone()

	rootCmd.AddCommand(debugknownCmd)
}
