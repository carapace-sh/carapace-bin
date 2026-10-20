// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fmtCmd = &cobra.Command{
	Use:   "fmt",
	Short: "Format mise TOML configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fmtCmd).Standalone()

	fmtCmd.Flags().BoolP("all", "a", false, "Format every config file mise loads, not only those in the current directory")
	fmtCmd.Flags().BoolP("check", "c", false, "Check whether the configs are formatted without rewriting them; exits 1 if any are not")
	fmtCmd.Flags().BoolP("help", "h", false, "Print help")
	fmtCmd.Flags().BoolP("stdin", "s", false, "Format TOML from stdin and print it to stdout, without touching any file")
	rootCmd.AddCommand(fmtCmd)
}
