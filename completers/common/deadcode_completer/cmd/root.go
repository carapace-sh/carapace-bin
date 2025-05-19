package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "deadcode",
	Short: "The deadcode command reports unreachable functions in Go programs",
	Long:  "https://pkg.go.dev/golang.org/x/tools/internal/cmd/deadcode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("cpuprofile", "cpuprofile", "", "write CPU profile to this file")
	rootCmd.Flags().StringS("filter", "filter", "", "report only packages matching this regular expression")
	rootCmd.Flags().BoolS("generated", "generated", false, "report dead functions in generated Go files")
	rootCmd.Flags().BoolS("line", "line", false, "show output in a line-oriented format")
	rootCmd.Flags().StringS("memprofile", "memprofile", "", "write memory profile to this file")
	rootCmd.Flags().StringS("tags", "tags", "", "comma-separated list of extra build tags")
	rootCmd.Flags().BoolS("test", "test", false, "include implicit test packages and executables")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cpuprofile": carapace.ActionFiles(),
		"filter":     golang.ActionPackages(),
		"memprofile": carapace.ActionFiles(),
		"tags":       golang.ActionBuildTags().UniqueList(","),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		golang.ActionPackages(),
	)
}
