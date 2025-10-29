package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_set_autoSelfUpdateCmd = &cobra.Command{
	Use:   "auto-self-update",
	Short: "The rustup auto self update mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_set_autoSelfUpdateCmd).Standalone()

	help_setCmd.AddCommand(help_set_autoSelfUpdateCmd)
}
