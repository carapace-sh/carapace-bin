package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var system_resetCmd = &cobra.Command{
	Use:   "reset [options]",
	Short: "Reset podman storage",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(system_resetCmd).Standalone()

	system_resetCmd.Flags().BoolP("force", "f", false, "Do not prompt for confirmation")
	systemCmd.AddCommand(system_resetCmd)
}
