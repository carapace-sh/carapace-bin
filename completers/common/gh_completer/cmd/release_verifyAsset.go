package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var release_verifyAssetCmd = &cobra.Command{
	Use:     "verify-asset [<tag>] <file-path>",
	Short:   "Verify that a given asset originated from a release",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(release_verifyAssetCmd).Standalone()

	release_verifyAssetCmd.Flags().String("custom-trusted-root", "", "Path to a trusted_root.jsonl file; likely for offline verification.")
	release_verifyAssetCmd.Flags().String("format", "", "Output format: {json}")
	release_verifyAssetCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	release_verifyAssetCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	release_verifyAssetCmd.Flag("custom-trusted-root").Hidden = true
	releaseCmd.AddCommand(release_verifyAssetCmd)

	carapace.Gen(release_verifyAssetCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("json"),
	})

	carapace.Gen(release_verifyAssetCmd).PositionalCompletion(
		action.ActionReleases(release_verifyAssetCmd),
		carapace.ActionFiles(),
	)
}
