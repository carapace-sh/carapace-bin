package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cover",
	Short: "analyze coverage profiles",
	Long:  "https://pkg.go.dev/cmd/cover",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute(opts ...func(cmd *cobra.Command)) error {
	for _, opt := range opts {
		opt(rootCmd)
	}
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()
	rootCmd.Flags().SetInterspersed(false)

	rootCmd.Flags().StringS("func", "func", "", "output coverage profile information for each function")
	rootCmd.Flags().StringS("html", "html", "", "generate HTML representation of coverage profile")
	rootCmd.Flags().StringS("mode", "mode", "", "coverage mode: set, count, atomic")
	rootCmd.Flags().StringS("o", "o", "", "file for output")
	rootCmd.Flags().StringS("outfilelist", "outfilelist", "", "file containing list of output files (one per line) if -pkgcfg is in use")
	rootCmd.Flags().StringS("pkgcfg", "pkgcfg", "", "enable full-package instrumentation mode using params from specified config file")
	rootCmd.Flags().StringS("var", "var", "", "name of coverage variable to generate (default \"GoCover\")")

	rootCmd.MarkFlagsMutuallyExclusive("html", "func", "mode")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"func":        carapace.ActionFiles(),
		"html":        carapace.ActionFiles(),
		"mode":        carapace.ActionValues("set", "count", "atomic"),
		"o":           carapace.ActionFiles(),
		"outfilelist": carapace.ActionFiles(),
		"pkgcfg":      carapace.ActionFiles(),
	})
}
