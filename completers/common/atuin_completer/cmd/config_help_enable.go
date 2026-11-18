package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var config_help_enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "Enable a feature, along with everything it depends on",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(config_help_enableCmd).Standalone()

	config_helpCmd.AddCommand(config_help_enableCmd)
}
