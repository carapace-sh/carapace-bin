package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/bazel_completer/cmd/common"
	"github.com/spf13/cobra"
)

var analyzeProfileCmd = &cobra.Command{
	Use:   "analyze-profile",
	Short: "Analyzes build profile data for the given profile data files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(analyzeProfileCmd).Standalone()
	common.AddRemoteCachingFlags(analyzeProfileCmd)
	common.AddBuildExecutionFlags(analyzeProfileCmd)
	common.AddToolchainFlags(analyzeProfileCmd)
	common.AddBuildTimeOptimizationFlags(analyzeProfileCmd)
	common.AddStarlarkSemanticFlags(analyzeProfileCmd)
	common.AddBzlmodFlags(analyzeProfileCmd)
	common.AddInputAlteringFlags(analyzeProfileCmd)
	common.AddInputValidationFlags(analyzeProfileCmd)
	common.AddLoggingFlags(analyzeProfileCmd)
	common.AddOutputFlags(analyzeProfileCmd)
	common.AddMiscellaneousFlags(analyzeProfileCmd)

	rootCmd.AddCommand(analyzeProfileCmd)

}
