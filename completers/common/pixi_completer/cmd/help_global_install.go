package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_global_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs the defined packages in a globally accessible location and exposes their command line applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_global_installCmd).Standalone()

	help_globalCmd.AddCommand(help_global_installCmd)
}
