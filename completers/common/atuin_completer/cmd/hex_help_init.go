package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hex_help_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Print shell code to initialize atuin-hex on shell startup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hex_help_initCmd).Standalone()

	hex_helpCmd.AddCommand(hex_help_initCmd)
}
