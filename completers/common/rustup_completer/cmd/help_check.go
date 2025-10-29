package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check for updates to Rust toolchains and rustup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_checkCmd).Standalone()

	helpCmd.AddCommand(help_checkCmd)
}
