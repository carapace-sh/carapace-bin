package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var coverageCmd = &cobra.Command{
	Use:   "coverage",
	Short: "Generate source-level coverage from SBF register traces",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(coverageCmd).Standalone()

	coverageCmd.Flags().BoolP("help", "h", false, "Print help")
	coverageCmd.Flags().String("output", "target/coverage/sbf.lcov", "Output path for the LCOV file")
	coverageCmd.Flags().Bool("skip-build", false, "Skip `cargo build-sbf`")
	coverageCmd.Flags().Bool("skip-run", false, "Skip the build+test phase and generate coverage from existing traces")
	coverageCmd.Flags().String("trace-dir", "target/coverage/traces", "Directory containing register trace files")
	rootCmd.AddCommand(coverageCmd)

	carapace.Gen(coverageCmd).FlagCompletion(carapace.ActionMap{
		"output":    carapace.ActionFiles(),
		"trace-dir": carapace.ActionDirectories(),
	})
}
