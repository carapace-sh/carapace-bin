package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check for updates to Rust toolchains and rustup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(checkCmd).Standalone()

	checkCmd.Flags().BoolP("help", "h", false, "Print help")
	checkCmd.Flags().Bool("no-self-update", false, "Don't check for self update when running the `rustup check` command")
	rootCmd.AddCommand(checkCmd)
}
