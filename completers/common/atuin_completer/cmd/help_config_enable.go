package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_config_enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable a feature, along with everything it depends on",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_config_enableCmd).Standalone()

	help_configCmd.AddCommand(help_config_enableCmd)
}
