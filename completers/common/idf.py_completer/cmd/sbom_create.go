package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sbomCreateCmd = &cobra.Command{
	Use:   "sbom-create",
	Short: "Create application SBOM in one of the supported formats",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sbomCreateCmd).Standalone()
	sbomCreateCmd.Flags().String("format", "", "Output SBOM format")
	sbomCreateCmd.Flags().String("output-file", "", "Output SBOM file")
	sbomCreateCmd.Flags().String("spdx-file", "", "Output SBOM file (deprecated alias for --output-file)")
	sbomCreateCmd.Flags().String("vex-format", "", "Vulnerability information format")
	sbomCreateCmd.Flags().String("vex-output", "", "Output VEX file path")
	sbomCreateCmd.Flag("spdx-file").Hidden = true
	rootCmd.AddCommand(sbomCreateCmd)
	carapace.Gen(sbomCreateCmd).FlagCompletion(carapace.ActionMap{
		"format":      carapace.ActionValues("cyclonedx-json", "spdx-json", "spdx-json-ld", "spdx-tag-value"),
		"output-file": carapace.ActionFiles(),
		"spdx-file":   carapace.ActionFiles(),
		"vex-format":  carapace.ActionValues("cyclonedx-json", "embed", "none", "openvex"),
		"vex-output":  carapace.ActionFiles(),
	})
}
