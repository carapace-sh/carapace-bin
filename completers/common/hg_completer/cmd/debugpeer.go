package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugpeerCmd = &cobra.Command{
	Use:    "debugpeer",
	Short:  "PATH",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugpeerCmd).Standalone()

	rootCmd.AddCommand(debugpeerCmd)
}
