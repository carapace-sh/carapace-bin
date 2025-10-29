package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var set_autoSelfUpdateCmd = &cobra.Command{
	Use:   "auto-self-update",
	Short: "The rustup auto self update mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(set_autoSelfUpdateCmd).Standalone()

	set_autoSelfUpdateCmd.Flags().BoolP("help", "h", false, "Print help")
	setCmd.AddCommand(set_autoSelfUpdateCmd)
}
