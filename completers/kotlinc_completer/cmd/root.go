package cmd

import (
	"os"
	"strings"

	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kotlinc",
	Short: "Kotlin command-line compiler",
	Long:  "https://kotlinlang.org/docs/command-line.html#manual-install",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	for _, arg := range os.Args {
		if strings.HasPrefix(arg, "-J") {
			name := strings.TrimPrefix(arg, "-")
			rootCmd.Flags().Bool(name, false, "")
		}
	}

	carapace.Override(carapace.Opts{
		LongShorthand: true,
	})

	return rootCmd.Execute()
}
func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().Bool("J", false, "Pass an option directly to JVM")
	rootCmd.Flags().String("P", "", "Pass an option to a plugin")
	rootCmd.Flags().Bool("Werror", false, "Report an error if there are any warnings")
	rootCmd.Flags().Bool("X", false, "Print a synopsis of advanced options")
	rootCmd.Flags().String("api-version", "", "Allow using declarations only from the specified version of bundled libraries")
	rootCmd.Flags().String("classpath", "", "List of directories and JAR/ZIP archives to search for user class files")
	rootCmd.Flags().StringS("d", "d", "", "Destination for generated class files")
	rootCmd.Flags().BoolP("expression", "e", false, "Evaluate the given string as a Kotlin script")
	rootCmd.Flags().BoolP("help", "h", false, "Print a synopsis of standard options")
	rootCmd.Flags().Bool("include-runtime", false, "Include Kotlin runtime into the resulting JAR")
	rootCmd.Flags().Bool("java-parameters", false, "Generate metadata for Java 1.8 reflection on method parameters")
	rootCmd.Flags().String("jdk-home", "", "Include a custom JDK from the specified location into the classpath")
	rootCmd.Flags().String("jvm-target", "", "Target version of the generated JVM bytecode")
	rootCmd.Flags().String("kotlin-home", "", "Path to the home directory of Kotlin compiler used for discovery of runtime libraries")
	rootCmd.Flags().String("language-version", "", "Provide source compatibility with the specified version of Kotlin")
	rootCmd.Flags().String("module-name", "", "Name of the generated .kotlin_module file")
	rootCmd.Flags().Bool("no-jdk", false, "Don't automatically include the Java runtime into the classpath")
	rootCmd.Flags().Bool("no-reflect", false, "Don't automatically include Kotlin reflection into the classpath")
	rootCmd.Flags().Bool("no-stdlib", false, "Don't automatically include the Kotlin/JVM stdlib and Kotlin reflection into the classpath")
	rootCmd.Flags().Bool("nowarn", false, "Generate no warnings")
	rootCmd.Flags().String("opt-in", "", "Enable usages of API that requires opt-in with an opt-in requirement marker with the given fully qualified name")
	rootCmd.Flags().Bool("progressive", false, "Enable progressive compiler mode.")
	rootCmd.Flags().Bool("script", false, "Evaluate the given Kotlin script (*.kts) file")
	rootCmd.Flags().String("script-templates", "", "Script definition template classes")
	rootCmd.Flags().Bool("verbose", false, "Enable verbose logging output")
	rootCmd.Flags().Bool("version", false, "Display compiler version")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"classpath":   carapace.ActionDirectories(),
		"d":           carapace.ActionDirectories(),
		"jdk-home":    carapace.ActionDirectories(),
		"jvm-target":  carapace.ActionValues("1.6", "1.8", "9", "10", "11", "12", "13", "14", "15", "16", "17", "18"),
		"kotlin-home": carapace.ActionDirectories(),
	})

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionFiles(".kt"),
	)
}
