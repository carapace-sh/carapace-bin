package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var attestation_trustedRootCmd = &cobra.Command{
	Use:   "trusted-root [--tuf-url <url> --tuf-root <file-path>] [--verify-only]",
	Short: "Output trusted_root.jsonl contents, likely for offline verification",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(attestation_trustedRootCmd).Standalone()

	attestation_trustedRootCmd.Flags().String("hostname", "", "Configure host to use")
	attestation_trustedRootCmd.Flags().String("tuf-root", "", "Path to the TUF root.json file on disk")
	attestation_trustedRootCmd.Flags().String("tuf-url", "", "URL to the TUF repository mirror")
	attestation_trustedRootCmd.Flags().Bool("verify-only", false, "Don't output trusted_root.jsonl contents")
	attestationCmd.AddCommand(attestation_trustedRootCmd)

	carapace.Gen(attestation_trustedRootCmd).FlagCompletion(carapace.ActionMap{
		"hostname": gh.ActionConfigHosts(),
		"tuf-root": carapace.ActionFiles(),
		// TODO tuf-url
	})
}
