package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/go_completer/cmd/common"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate Go files by processing source",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(generateCmd).Standalone()
	generateCmd.Flags().SetInterspersed(false)

	generateCmd.Flags().StringS("run", "run", "", "specifies a regular expression to select matching directives")
	generateCmd.Flags().StringS("skip", "skip", "", "specifies a regular expression to suppress matching directives")
	common.AddPackageBuildFlags(generateCmd)
	rootCmd.AddCommand(generateCmd)

	carapace.Gen(generateCmd).FlagCompletion(carapace.ActionMap{
		"C":       carapace.ActionDirectories(),
		"overlay": carapace.ActionFiles(".json"),
	})

	carapace.Gen(generateCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".go"),
	)
}
