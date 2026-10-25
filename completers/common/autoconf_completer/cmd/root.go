package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "autoconf",
	Short: "Generate a configuration script from a TEMPLATE-FILE",
	Long:  "https://linux.die.net/man/1/autoconf",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolP("debug", "d", false, "don't remove temporary files")
	rootCmd.Flags().BoolP("force", "f", false, "consider all files obsolete")
	rootCmd.Flags().BoolP("help", "h", false, "print this help, then exit")
	rootCmd.Flags().StringP("include", "I", "", "append directory DIR to search path")
	rootCmd.Flags().BoolP("initialization", "i", false, "also trace Autoconf's initialization process")
	rootCmd.Flags().StringP("output", "o", "", "save output in FILE (stdout is the default)")
	rootCmd.Flags().StringP("prepend-include", "B", "", "prepend directory DIR to search path")
	rootCmd.Flags().StringP("trace", "t", "", "report the list of calls to MACRO")
	rootCmd.Flags().BoolP("verbose", "v", false, "verbosely report processing")
	rootCmd.Flags().BoolP("version", "V", false, "print version number, then exit")
	rootCmd.Flags().StringP("warnings", "W", "", "report the warnings falling in CATEGORY")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"include":         carapace.ActionDirectories(),
		"output":          carapace.ActionFiles(),
		"prepend-include": carapace.ActionDirectories(),
		"warnings": carapace.ActionValuesDescribed(
			"all", "all the warnings",
			"cross", "cross compilation issues",
			"extra-portability", "extra portability issues related to obscure tools",
			"gnu", "GNU coding standards (default in gnu and gnits modes)",
			"no-CATEGORY", "turn off warnings in CATEGORY",
			"none", "turn off all the warnings",
			"obsolete", "obsolete features or constructions (default)",
			"override", "user redefinitions of Automake rules or variables",
			"portability", "portability issues (default in gnu and gnits modes)",
			"portability-recursive", "nested Make variables (default with -Wportability)",
			"syntax", "dubious syntactic constructs (default)",
			"unsupported", "unsupported or incomplete features (default)",
		),
	})

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
