package cmd

import (
	"os"
	"strings"

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

	rootCmd.Flags().StringS("D", "D", "", "Define a macro")
	rootCmd.Flags().BoolS("E", "E", false, "Preprocess only; do not compile, assemble or link")
	rootCmd.Flags().StringS("I", "I", "", "Add dir to the end of the list of include search paths")
	rootCmd.Flags().StringS("L", "L", "", "Add directory to library search path")
	rootCmd.Flags().BoolS("S", "S", false, "Stop after the stage of compilation proper; do not assemble")
	rootCmd.Flags().StringS("U", "U", "", "Cancel a macro")
	rootCmd.Flags().Bool("Wall", false, "Enable most warning messages")
	rootCmd.Flags().Bool("Werror", false, "Make all warnings into errors")
	rootCmd.Flags().Bool("Wextra", false, "Enable extra warning messages")
	rootCmd.Flags().BoolS("c", "c", false, "Compile or assemble the source files, but do not link")
	rootCmd.Flags().String("completion", "", "Provide bash completion for options starting with provided string")
	rootCmd.Flags().Bool("fPIC", false, "Generate position-independent code for shared libraries")
	rootCmd.Flags().Bool("fPIE", false, "Generate position-independent code for executables")
	rootCmd.Flags().Bool("fpic", false, "Similar to -fPIC, but smaller")
	rootCmd.Flags().Bool("fpie", false, "Similar to -fPIE, but smaller")
	rootCmd.Flags().String("help", "", "Display help information")
	rootCmd.Flags().String("imacros", "", "Like -include, but with the macro expansion")
	rootCmd.Flags().String("include", "", "Process file as if #include \"file\" appeared as the first line of the primary source file")
	rootCmd.Flags().StringS("l", "l", "", "Search for library LIBNAME")
	rootCmd.Flags().Bool("m32", false, "Generate 32bit i386 code")
	rootCmd.Flags().Bool("m64", false, "Generate 64bit x86-64 code")
	rootCmd.Flags().String("march", "", "Generate code for given CPU")
	rootCmd.Flags().String("mtune", "", "Tune code for given CPU")
	rootCmd.Flags().Bool("no-pie", false, "Don't create a position independent executable")
	rootCmd.Flags().StringS("o", "o", "", "Place output into <file>")
	rootCmd.Flags().String("param", "", "Specify a parameter for the compiler")
	rootCmd.Flags().Bool("pedantic", false, "Issue warnings needed by strict compliance to the standard")
	rootCmd.Flags().Bool("pedantic-errors", false, "Like -pedantic but issue errors instead of warnings")
	rootCmd.Flags().Bool("pie", false, "Create a position independent executable")
	rootCmd.Flags().Bool("pipe", false, "Use pipes rather than intermediate files")
	rootCmd.Flags().Bool("pthread", false, "Use the pthreads library")
	rootCmd.Flags().Bool("save-temps", false, "Do not delete intermediate files")
	rootCmd.Flags().Bool("shared", false, "Create a shared library")
	rootCmd.Flags().Bool("static", false, "Do not link against shared libraries")
	rootCmd.Flags().String("std", "", "Assume that the input sources are for <standard>")
	rootCmd.Flags().Bool("target-help", false, "Display target specific command line options")
	rootCmd.Flags().BoolS("v", "v", false, "Enable verbose output")
	rootCmd.Flags().Bool("version", false, "Display the compiler's version")
	rootCmd.Flags().Bool("w", false, "Inhibit all warning messages")
	rootCmd.Flags().StringS("x", "x", "", "Specify the language of the following input files")

	rootCmd.Flag("help").NoOptDefVal = " "
	rootCmd.Flag("completion").NoOptDefVal = " "

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"D":          carapace.ActionValues(),
		"I":          carapace.ActionDirectories(),
		"L":          carapace.ActionDirectories(),
		"U":          carapace.ActionValues(),
		"completion": carapace.ActionValues("common", "optimizers", "warnings", "target", "params", "undocumented"),
		"help":       carapace.ActionValues("common", "optimizers", "warnings", "target", "params", "undocumented"),
		"imacros":    carapace.ActionFiles(),
		"include":    carapace.ActionFiles(),
		"l":          carapace.ActionValues(),
		"march":      action.ActionCPUs(),
		"mtune":      action.ActionCPUs(),
		"o":          carapace.ActionFiles(),
		"param":      action.ActionParams(),
		"std":        action.ActionStandards(),
		"x":          action.ActionLanguages(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if strings.HasPrefix(c.Value, "-") {
				return carapace.ActionValues()
			}
			return carapace.Batch(
				carapace.ActionFiles(".c", ".cc", ".cpp", ".cxx", ".C", ".h", ".hpp", ".hxx", ".s", ".S", ".i", ".ii", ".o", ".a", ".so"),
				carapace.ActionFiles().NoSpace(),
			).ToA()
		}),
	)

	carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
		os.Setenv("CARAPACE_LENIENT", "1")
	})
}