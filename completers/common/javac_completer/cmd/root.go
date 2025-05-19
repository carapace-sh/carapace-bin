package cmd

import (
	"os"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/javac_completer/cmd/action"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "javac",
	Short: "Reads Java class and interface definitions and compiles them into bytecode and class files",
	Long:  "https://en.wikipedia.org/wiki/Javac",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("?", "?", false, "Print help message")
	rootCmd.Flags().BoolS("A", "A", false, "Options to pass to annotation processors")
	rootCmd.Flags().BoolS("J", "J", false, "Pass flag directly to the runtime system")
	rootCmd.Flags().BoolS("Werror", "Werror", false, "Terminate compilation if warnings occur")
	rootCmd.Flags().String("add-modules", "", "Root modules to resolve in addition to the initial modules")
	rootCmd.Flags().StringP("boot-class-path", "bootclasspath", "", "Override location of bootstrap class files")
	rootCmd.Flags().StringP("class-path", "classpath", "", "Specify where to find user class files and annotation processors")
	rootCmd.Flags().StringS("cp", "cp", "", "Specify where to find user class files and annotation processors")
	rootCmd.Flags().BoolS("d", "d", false, "Specify where to place generated class files")
	rootCmd.Flags().BoolS("deprecation", "deprecation", false, "Output source locations where deprecated APIs are used")
	rootCmd.Flags().Bool("enable-preview", false, "Enable preview languge features")
	rootCmd.Flags().StringS("encoding", "encoding", "", "Specify character encoding used by source files")
	rootCmd.Flags().StringS("endorseddirs", "endorseddirs", "", "Override location of endorsed standards path")
	rootCmd.Flags().StringS("extdirs", "extdirs", "", "Override location of installed extensions")
	rootCmd.Flags().StringS("g", "g", "", "Generates all debugging information")
	rootCmd.Flags().BoolS("h", "h", false, "Specify where to place generated native header files")
	rootCmd.Flags().BoolS("help", "help", false, "Print help message")
	rootCmd.Flags().BoolS("help-extra", "X", false, "Print help on extra options")
	rootCmd.Flags().StringS("implicit", "implicit", "", "Specify whether or not to generate class files for implicitly referenced files")
	rootCmd.Flags().String("limit-modules", "", "Limit the universe of observable modules")
	rootCmd.Flags().String("module", "", "Compile only the specified module")
	rootCmd.Flags().StringP("module-path", "p", "", "Root modules to resolve in addition to the initial modules")
	rootCmd.Flags().String("module-source-path", "", "Specify where to find input source files for multiple modules")
	rootCmd.Flags().String("module-version", "", "Specify version of modules that are being compiled")
	rootCmd.Flags().BoolS("nowarn", "nowarn", false, "Generate no warnings")
	rootCmd.Flags().BoolS("parameters", "parameters", false, "Generate metadata for reflection on method parameters")
	rootCmd.Flags().StringS("proc", "proc", "", "Control whether annotation processing and/or compilation is done.")
	rootCmd.Flags().StringS("processor", "processor", "", "Names of the annotation processors to run; bypasses default discovery process")
	rootCmd.Flags().String("processor-module-path", "", "Specify a module path where to find annotation processors")
	rootCmd.Flags().StringP("processor-path", "processorpath", "", "Specify where to find annotation processors")
	rootCmd.Flags().StringS("profile", "profile", "", "Check that API used is available in the specified profile")
	rootCmd.Flags().String("release", "", "Compile for a specific VM version")
	rootCmd.Flags().BoolS("s", "s", false, "Specify where to place generated source files")
	rootCmd.Flags().StringS("source", "source", "", "Provide source compatibility with specified release")
	rootCmd.Flags().StringP("source-path", "sourcepath", "", "Specify where to find input source files")
	rootCmd.Flags().String("system", "", "Override location of system modules")
	rootCmd.Flags().StringS("target", "target", "", "Generate class files for specific VM version")
	rootCmd.Flags().String("upgrade-module-path", "", "Override location of upgradeable modules")
	rootCmd.Flags().BoolS("verbose", "verbose", false, "Output messages about what the compiler is doing")
	rootCmd.Flags().BoolP("version", "version", false, "Version information")

	rootCmd.Flag("g").NoOptDefVal = " "
	rootCmd.Flag("g").OptargDelimiter = ':'
	rootCmd.Flag("implicit").NoOptDefVal = " "
	rootCmd.Flag("implicit").OptargDelimiter = ':'
	rootCmd.Flag("proc").NoOptDefVal = " "
	rootCmd.Flag("proc").OptargDelimiter = ':'

	rootCmd.Flag("A").NoOptDefVal = " "
	rootCmd.Flag("A").OptargDelimiter = -1
	rootCmd.Flag("J").NoOptDefVal = " "
	rootCmd.Flag("J").OptargDelimiter = -1

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"boot-class-path": carapace.ActionFiles(".jar", ".zip").UniqueList(":"),
		"class-path":      carapace.ActionFiles(".jar", ".zip").UniqueList(":"),
		"cp":              carapace.ActionFiles(".jar", ".zip").UniqueList(":"),
		"d":               carapace.ActionDirectories(),
		// TODO encoding
		"endorseddirs": carapace.ActionDirectories(),
		"extdirs":      carapace.ActionDirectories(),
		"g": carapace.ActionValuesDescribed(
			"none", "Does not generate any debugging information",
			"source", "Source file debugging information",
			"lines", "Line number debugging information",
			"vars", "Local variable debugging information",
		).UniqueList(","),
		"h":                     carapace.ActionDirectories(),
		"implicit":              carapace.ActionValues("none", "class"),
		"module-path":           carapace.ActionDirectories(),
		"module-source-path":    carapace.ActionDirectories(),
		"proc":                  carapace.ActionValues("none", "only"),
		"processor-module-path": carapace.ActionDirectories(),
		"processor-path":        carapace.ActionDirectories(),
		"profile":               carapace.ActionValues("compact1", "compact2", "compact3"),
		"s":                     carapace.ActionDirectories(),
		"source":                action.ActionReleases(),
		"source-path":           carapace.ActionDirectories(),
		"system":                carapace.ActionValues("jdk", "none"),
		"target":                action.ActionReleases(), // TODO limit to `>+ source version`
		"upgrade-module-path":   carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if strings.HasPrefix(c.Value, "@") {
				c.Value = strings.TrimPrefix(c.Value, "@")
				return carapace.ActionFiles().Invoke(c).Prefix("@").ToA()
			}
			return carapace.ActionFiles(".java")
		}),
	)

	carapace.Gen(rootCmd).PreRun(func(cmd *cobra.Command, args []string) {
		os.Setenv("CARAPACE_LENIENT", "1") // TODO hacky -A and -J support for now
	})
}
