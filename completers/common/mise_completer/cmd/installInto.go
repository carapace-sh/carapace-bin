// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var installIntoCmd = &cobra.Command{
	Use:   "install-into",
	Short: "Install a tool version into a directory outside mise",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(installIntoCmd).Standalone()

	installIntoCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(installIntoCmd)

	carapace.Gen(installIntoCmd).PositionalCompletion(
		mise.ActionToolVersions(),
	)
}
