package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_packages_brewCmd = &cobra.Command{
	Use:   "brew",
	Short: "Manage Homebrew tap URLs in `[bootstrap.brew.taps]`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_packages_brewCmd).Standalone()

	bootstrap_packages_brewCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrap_packagesCmd.AddCommand(bootstrap_packages_brewCmd)
}
