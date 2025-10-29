package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var set_help_autoSelfUpdateCmd = &cobra.Command{
	Use:   "auto-self-update",
	Short: "The rustup auto self update mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(set_help_autoSelfUpdateCmd).Standalone()

	set_helpCmd.AddCommand(set_help_autoSelfUpdateCmd)
}
