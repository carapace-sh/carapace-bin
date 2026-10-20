// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_packages_exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Print the configured `nix:` packages as a NixOS module",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_packages_exportCmd).Standalone()

	bootstrap_packages_exportCmd.Flags().String("format", "", "Output format")
	bootstrap_packages_exportCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_packages_exportCmd.MarkFlagRequired("format")
	bootstrap_packagesCmd.AddCommand(bootstrap_packages_exportCmd)

	carapace.Gen(bootstrap_packages_exportCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("nix"),
	})
}
