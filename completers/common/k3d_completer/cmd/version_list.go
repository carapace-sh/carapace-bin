package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var version_listCmd = &cobra.Command{
	Use:     "list COMPONENT",
	Short:   "List k3d/K3s versions. Component can be one of 'k3d', 'k3s', 'k3d-proxy', 'k3d-tools'.",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(version_listCmd).Standalone()

	version_listCmd.Flags().StringP("exclude", "e", "", "Exclude Regexp (default excludes pre-releases, arch-specific tags and digests)")
	version_listCmd.Flags().StringP("format", "f", "", "[DEPRECATED] Use --output instead")
	version_listCmd.Flags().StringP("include", "i", "", "Include Regexp (default includes everything")
	version_listCmd.Flags().StringP("limit", "l", "", "Limit number of tags in output (0 = unlimited)")
	version_listCmd.Flags().StringP("output", "o", "", "Output Format [raw | repo]")
	version_listCmd.Flags().StringP("sort", "s", "", "Sort Mode (asc | desc | off)")
	versionCmd.AddCommand(version_listCmd)

	carapace.Gen(version_listCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionValues("raw", "repo"),
		"sort":   carapace.ActionValues("asc", "desc", "off"),
	})

	carapace.Gen(version_listCmd).PositionalCompletion(
		carapace.ActionValues("k3d", "k3s", "k3d-proxy", "k3d-tools"),
	)
}
