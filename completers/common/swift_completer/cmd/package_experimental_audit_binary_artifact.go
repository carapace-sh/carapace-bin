package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var package_experimentalAuditBinaryArtifactCmd = &cobra.Command{
	Use:   "experimental-audit-binary-artifact",
	Short: "Audit a static library binary artifact for undefined symbols",
	Args:  cobra.ExactArgs(1),
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(package_experimentalAuditBinaryArtifactCmd).Standalone()
	package_experimentalAuditBinaryArtifactCmd.Flags().SetInterspersed(false)

	package_experimentalAuditBinaryArtifactCmd.Flags().BoolP("help", "h", false, "Show help information")
	package_experimentalAuditBinaryArtifactCmd.Flags().Bool("version", false, "Show the version")

	packageCmd.AddCommand(package_experimentalAuditBinaryArtifactCmd)

	carapace.Gen(package_experimentalAuditBinaryArtifactCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
