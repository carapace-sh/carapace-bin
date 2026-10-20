// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var linkCmd = &cobra.Command{
	Use:     "link",
	Short:   "Symlink a tool version into mise",
	Aliases: []string{"ln"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(linkCmd).Standalone()

	linkCmd.Flags().BoolP("force", "f", false, "Replace an existing installation of this version")
	linkCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(linkCmd)

	carapace.Gen(linkCmd).PositionalCompletion(
		mise.ActionToolVersions(),
		carapace.ActionDirectories(),
	)
}
