package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_hex_initCmd = &cobra.Command{
	Use:   "init",
	Short: "Print shell code to initialize atuin-hex on shell startup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_hex_initCmd).Standalone()

	help_hexCmd.AddCommand(help_hex_initCmd)
}
