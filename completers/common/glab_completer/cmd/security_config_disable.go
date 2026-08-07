package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var security_config_disableCmd = &cobra.Command{
	Use:   "disable <profile> [flags]",
	Short: "Disable a security scan profile for a project. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(security_config_disableCmd).Standalone()

	security_configCmd.AddCommand(security_config_disableCmd)
}
