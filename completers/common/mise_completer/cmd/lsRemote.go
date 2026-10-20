// Generated with `carapace --codegen` from the usage spec of mise 2026.10.15 (`mise usage`).
// Completions below the flag definitions were wired by hand.

package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/mise"
	"github.com/spf13/cobra"
)

var lsRemoteCmd = &cobra.Command{
	Use:     "ls-remote",
	Short:   "List tool versions available to install",
	Aliases: []string{"list-all", "list-remote"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lsRemoteCmd).Standalone()

	lsRemoteCmd.Flags().Bool("all", false, "List available versions for every tool mise currently knows about")
	lsRemoteCmd.Flags().String("before", "", "Only show versions released before a date or at least a duration ago")
	lsRemoteCmd.Flags().BoolP("help", "h", false, "Print help")
	lsRemoteCmd.Flags().BoolP("json", "J", false, "Output in JSON format, with metadata such as `created_at` when available")
	lsRemoteCmd.Flags().String("minimum-release-age", "", "Only show versions released before a date or at least a duration ago")
	lsRemoteCmd.Flags().Bool("no-versions-host", false, "Fetch versions from the tool's upstream instead of mise's version mirror")
	lsRemoteCmd.Flags().Bool("prerelease", false, "Include prerelease versions")
	lsRemoteCmd.Flags().Bool("strict-metadata", false, "Exit with an error when release metadata cannot be fetched")
	lsRemoteCmd.Flag("before").Hidden = true
	rootCmd.AddCommand(lsRemoteCmd)

	carapace.Gen(lsRemoteCmd).PositionalCompletion(
		mise.ActionToolVersions(),
	)
}
