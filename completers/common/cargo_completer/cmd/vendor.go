package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vendorCmd = &cobra.Command{
	Use:   "vendor",
	Short: "Vendor all dependencies for a project locally",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vendorCmd).Standalone()

	vendorCmd.Flags().BoolP("help", "h", false, "Print help")
	vendorCmd.Flags().String("lockfile-path", "", "Path to Cargo.lock (unstable)")
	vendorCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	vendorCmd.Flags().Bool("no-delete", false, "Don't delete older crates in the vendor directory")
	vendorCmd.Flags().Bool("respect-source-config", false, "Respect `[source]` config in `.cargo/config`")
	vendorCmd.Flags().StringSliceP("sync", "s", nil, "Additional `Cargo.toml` to sync and vendor")
	vendorCmd.Flags().Bool("versioned-dirs", false, "Always include version in subdir name")
	rootCmd.AddCommand(vendorCmd)

	carapace.Gen(vendorCmd).FlagCompletion(carapace.ActionMap{
		"lockfile-path": carapace.ActionFiles(),
		"manifest-path": carapace.ActionFiles(),
		"sync":          carapace.ActionFiles().List(","),
	})
}
