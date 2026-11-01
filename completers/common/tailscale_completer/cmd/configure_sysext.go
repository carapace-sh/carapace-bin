package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configure_sysextCmd = &cobra.Command{
	Use:   "sysext",
	Short: "Manage the system extension for macOS (Standalone variant)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configure_sysextCmd).Standalone()

	configureCmd.AddCommand(configure_sysextCmd)
}
