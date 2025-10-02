package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/spf13/cobra"
)

var release_verifyCmd = &cobra.Command{
	Use:     "verify [<tag>]",
	Short:   "Verify the attestation for a release",
	GroupID: "Targeted commands",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(release_verifyCmd).Standalone()

	release_verifyCmd.Flags().String("custom-trusted-root", "", "Path to a trusted_root.jsonl file; likely for offline verification.")
	release_verifyCmd.Flags().String("format", "", "Output format: {json}")
	release_verifyCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	release_verifyCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	release_verifyCmd.Flag("custom-trusted-root").Hidden = true
	releaseCmd.AddCommand(release_verifyCmd)

	carapace.Gen(release_verifyCmd).FlagCompletion(carapace.ActionMap{
		"custom-trusted-root": carapace.ActionFiles(),
		"format":              carapace.ActionValues("json"),
	})

	carapace.Gen(release_verifyCmd).PositionalCompletion(
		action.ActionReleases(release_verifyCmd),
	)
}
