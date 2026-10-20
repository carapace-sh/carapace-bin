// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_macosDefaultsCmd = &cobra.Command{
	Use:    "macos-defaults",
	Short:  "Manage macOS preferences (use `mise bootstrap macos defaults`)",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_macosDefaultsCmd).Standalone()

	bootstrap_macosDefaultsCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrapCmd.AddCommand(bootstrap_macosDefaultsCmd)
}
