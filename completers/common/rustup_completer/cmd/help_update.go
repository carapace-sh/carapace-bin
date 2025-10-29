package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update Rust toolchains and rustup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_updateCmd).Standalone()

	helpCmd.AddCommand(help_updateCmd)
}
