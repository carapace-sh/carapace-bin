// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var globalCmd = &cobra.Command{
	Use:    "global",
	Short:  "Set or show the global tool versions (deprecated; use `mise use -g`)",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalCmd).Standalone()

	globalCmd.Flags().Bool("fuzzy", false, "Save the version as written, such as `24`")
	globalCmd.Flags().BoolP("help", "h", false, "Print help")
	globalCmd.Flags().Bool("path", false, "Print the path of the global config file")
	globalCmd.Flags().Bool("pin", false, "Save the exact version, such as `24.11.0`")
	globalCmd.Flags().StringSlice("remove", nil, "Remove these tools from the global config")
	globalCmd.Flags().StringSlice("rm", nil, "Remove these tools from the global config")
	globalCmd.Flags().StringSlice("unset", nil, "Remove these tools from the global config")
	globalCmd.Flag("rm").Hidden = true
	globalCmd.Flag("unset").Hidden = true
	rootCmd.AddCommand(globalCmd)

	carapace.Gen(globalCmd).PositionalAnyCompletion(
		mise.ActionToolVersions(),
	)
}
