package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fix",
	Long:  "https://pkg.go.dev/cmd/fix",
	Short: "Fix finds Go programs that use old APIs",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("diff", "diff", false, "display diffs instead of rewriting files")
	rootCmd.Flags().StringS("force", "force", "", "force these fixes to run even if the code looks updated")
	rootCmd.Flags().StringS("go", "go", "", "go language version for files")
	rootCmd.Flags().StringS("r", "r", "", "restrict the rewrites to this comma-separated list")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"force": golang.ActionRewrites().UniqueList(","),
		"go":    golang.ActionVersions(),
		"r":     golang.ActionRewrites().UniqueList(","),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
