package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/gh"
	"github.com/spf13/cobra"
)

var attestation_inspectCmd = &cobra.Command{
	Use:    "inspect <path-to-sigstore-bundle>",
	Short:  "Inspect a Sigstore bundle",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(attestation_inspectCmd).Standalone()

	attestation_inspectCmd.Flags().String("format", "", "Output format: {json}")
	attestation_inspectCmd.Flags().String("hostname", "", "Configure host to use")
	attestation_inspectCmd.Flags().StringP("jq", "q", "", "Filter JSON output using a jq `expression`")
	attestation_inspectCmd.Flags().StringP("template", "t", "", "Format JSON output using a Go template; see \"gh help formatting\"")
	attestationCmd.AddCommand(attestation_inspectCmd)

	carapace.Gen(attestation_inspectCmd).FlagCompletion(carapace.ActionMap{
		"format":   carapace.ActionValues("json"),
		"hostname": gh.ActionConfigHosts(),
	})

	carapace.Gen(attestation_inspectCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
