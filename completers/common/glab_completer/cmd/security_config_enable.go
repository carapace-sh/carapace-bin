package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var security_config_enableCmd = &cobra.Command{
	Use:   "enable <profile> [flags]",
	Short: "Enable a security scan profile for a project. (EXPERIMENTAL)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(security_config_enableCmd).Standalone()

	security_configCmd.AddCommand(security_config_enableCmd)
}
