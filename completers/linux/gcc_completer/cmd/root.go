package cmd

import (
	"os"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/gcc_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gcc",
	Short: "GNU project C and C++ compiler",
	Long:  "https://gcc.gnu.org/",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("E", "E", false, "Preprocess only; do not compile, assemble or link")
	rootCmd.Flags().BoolS("S", "S", false, "Compile only; do not assemble or link")
	rootCmd.Flags().BoolS("c", "c", false, "Compile and assemble, but do not link")
	rootCmd.Flags().StringS("o", "o", "", "Place output into <file>")
	rootCmd.Flags().Bool("pedantic", false, "Issue warnings needed by strict compliance to the standard")
	rootCmd.Flags().Bool("pedantic-errors", false, "Like -pedantic but issue errors instead of warnings")
	rootCmd.Flags().Bool("Wall", false, "Enable most warning messages")
	rootCmd.Flags().Bool("Wextra", false, "Enable extra warning messages")
	rootCmd.Flags().Bool("Werror", false, "Make all warnings into errors")
	rootCmd.Flags().Bool("w", false, "Inhibit all warning messages")
	rootCmd.Flags().Bool("ansi", false, "A synonym for -std=c89")
	rootCmd.Flags().Bool("pthread", false, "Use the pthreads library")
	rootCmd.Flags().Bool("pipe", false, "Use pipes rather than intermediate files")
	rootCmd.Flags().Bool("shared", false, "Create a shared library")
	rootCmd.Flags().Bool("static", false, "Do not link against shared libraries")
	rootCmd.Flags().Bool("pie", false, "Create a position independent executable")
	rootCmd.Flags().Bool("no-pie", false, "Don't create a position independent executable")
	rootCmd.Flags().Bool("fPIC", false, "Generate position-independent code for shared libraries")
	rootCmd.Flags().Bool("fPIE", false, "Generate position-independent code for executables")
	rootCmd.Flags().Bool("fpic", false, "Similar to -fPIC, but smaller")
	rootCmd.Flags().Bool("fpie", false, "Similar to -fPIE, but smaller")
	rootCmd.Flags().Bool("m32", false, "Generate 32bit i386 code")
	rootCmd.Flags().Bool("m64", false, "Generate 64bit x86-64 code")
	rootCmd.Flags().StringS("I", "I", "", "Add dir to the end of the list of include search paths")
	rootCmd.Flags().StringS("L", "L", "", "Add directory to library search path")
	rootCmd.Flags().StringS("l", "l", "", "Search for library LIBNAME")
	rootCmd.Flags().StringS("D", "D", "", "Define a macro")
	rootCmd.Flags().StringS("U", "U", "", "Cancel a macro")
	rootCmd.Flags().String("include", "", "Process file as if #include \"file\" appeared as the first line of the primary source file")
	rootCmd.Flags().String("imacros", "", "Like -include, but with the macro expansion")
	rootCmd.Flags().String("save-temps", "", "Do not delete intermediate files")
	rootCmd.Flags().String("std", "", "Assume that the input sources are for <standard>")
	rootCmd.Flags().String("march", "", "Generate code for given CPU")
	rootCmd.Flags().String("mtune", "", "Tune code for given CPU")
	rootCmd.Flags().String("x", "", "Specify the language of the following input files")

	rootCmd.Flag("save-temps").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"D":          action.ActionFlagValues("-D"),
		"I":          carapace.ActionDirectories(),
		"L":          carapace.ActionDirectories(),
		"U":          action.ActionFlagValues("-U"),
		"include":    carapace.ActionFiles(),
		"imacros":    carapace.ActionFiles(),
		"l":          carapace.ActionValues(),
		"march":      action.ActionFlagValues("-march"),
		"mtune":      action.ActionFlagValues("-mtune"),
		"o":          carapace.ActionFiles(),
		"save-temps": action.ActionFlagValues("-save-temps"),
		"std":        action.ActionFlagValues("-std"),
		"x":          action.ActionFlagValues("-x"),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return action.ActionDynamic("")
		}),
	)

	carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
		os.Setenv("CARAPACE_LENIENT", "1")
	})
}