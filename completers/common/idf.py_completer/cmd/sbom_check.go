package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sbomCheckCmd = &cobra.Command{
	Use:   "sbom-check",
	Short: "Check application SBOM file or path for vulnerabilities",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sbomCheckCmd).Standalone()
	sbomCheckCmd.Flags().String("file", "", "SBOM file to check")
	sbomCheckCmd.Flags().String("format", "", "Report format")
	sbomCheckCmd.Flags().Bool("no-sync-db", false, "Prevent automatic update of local NVD database")
	sbomCheckCmd.Flags().Bool("nvd-api", false, "Use NVD REST API for vulnerabilities scan")
	sbomCheckCmd.Flags().String("path", "", "Path to recursively search for SBOM YAML manifest files")
	sbomCheckCmd.Flags().String("report-file", "", "Report output file")
	sbomCheckCmd.Flags().String("spdx-file", "", "SBOM file to check (deprecated alias for --file)")
	sbomCheckCmd.Flag("spdx-file").Hidden = true
	carapace.Gen(sbomCheckCmd).FlagCompletion(carapace.ActionMap{
		"file":        carapace.ActionFiles(),
		"format":      carapace.ActionValues("csv", "json", "markdown", "table"),
		"path":        carapace.ActionDirectories(),
		"report-file": carapace.ActionFiles(),
		"spdx-file":   carapace.ActionFiles(),
	})
	rootCmd.AddCommand(sbomCheckCmd)
}
