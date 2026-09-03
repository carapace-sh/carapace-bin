package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_snapshotCmd = &cobra.Command{
	Use:    "snapshot",
	Short:  "[DEPRECATED] Trigger a snapshot in the op log",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_snapshotCmd).Standalone()

	debug_snapshotCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	debugCmd.AddCommand(debug_snapshotCmd)
}