package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_help_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs the defined packages in a globally accessible location and exposes their command line applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_help_installCmd).Standalone()

	global_helpCmd.AddCommand(global_help_installCmd)
}
