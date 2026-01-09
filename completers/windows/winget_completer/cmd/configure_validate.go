package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var configure_validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validates a configuration file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(configure_validateCmd).Standalone()

	configureCmd.AddCommand(configure_validateCmd)
}
