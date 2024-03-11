package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/golang"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "goimports",
	Short: "updates your Go import lines",
	Long:  "https://pkg.go.dev/golang.org/x/tools/cmd/goimports",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("cpuprofile", "cpuprofile", "", "CPU profile output")
	rootCmd.Flags().BoolS("d", "d", false, "display diffs instead of rewriting files")
	rootCmd.Flags().BoolS("e", "e", false, "report all errors")
	rootCmd.Flags().BoolS("format-only", "format-only", false, "don't fix imports and only format")
	rootCmd.Flags().BoolS("l", "l", false, "list files whose formatting differs from goimport's")
	rootCmd.Flags().StringS("local", "local", "", "put imports beginning with this string after 3rd-party packages")
	rootCmd.Flags().StringS("memprofile", "memprofile", "", "memory profile output")
	rootCmd.Flags().StringS("memrate", "memrate", "", "if > 0, sets runtime.MemProfileRate")
	rootCmd.Flags().StringS("srcdir", "srcdir", "", "choose imports as if source code is from dir")
	rootCmd.Flags().StringS("trace", "trace", "", "trace profile output")
	rootCmd.Flags().BoolS("v", "v", false, "verbose logging")
	rootCmd.Flags().BoolS("w", "w", false, "write result to (source) file instead of stdout")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"cpuprofile": carapace.ActionFiles(),
		"local":      golang.ActionPackages().List(","),
		"memprofile": carapace.ActionFiles(),
		"srcdir":     carapace.ActionDirectories(),
		"trace":      carapace.ActionFiles(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
