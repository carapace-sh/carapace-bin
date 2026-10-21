package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var analyzeCmd = &cobra.Command{
	Use:   "analyze",
	Short: "Analyze the project's Dart code",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(analyzeCmd).Standalone()

	analyzeCmd.Flags().Bool("android", false, "Analyze Android sub-project. Used by internal tools only.")
	analyzeCmd.Flags().Bool("benchmark", false, "Also output the analysis time.")
	analyzeCmd.Flags().String("build-variant", "", "Sets the Android build variant to be analyzed.")
	analyzeCmd.Flags().String("configuration", "", "Sets the iOS build configuration to be analyzed.")
	analyzeCmd.Flags().Bool("congratulate", false, "Show output even when there are no errors, warnings, hints, or lints. Ignored if \"--watch\" is specified.")
	analyzeCmd.Flags().Bool("current-package", false, "Analyze the current project, if applicable.")
	analyzeCmd.Flags().String("dart-sdk", "", "The path to the Dart SDK.")
	analyzeCmd.Flags().Bool("dartdocs", false, "(deprecated) List every public member that is lacking documentation. This command will be removed in a future version of Flutter.")
	analyzeCmd.Flags().Bool("fatal-infos", false, "Treat info level issues as fatal.")
	analyzeCmd.Flags().Bool("fatal-warnings", false, "Treat warning level issues as fatal.")
	analyzeCmd.Flags().Bool("flutter-repo", false, "Include all the examples and tests from the Flutter repository.")
	analyzeCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	analyzeCmd.Flags().Bool("ios", false, "Analyze iOS Xcode sub-project. Used by internal tools only.")
	analyzeCmd.Flags().Bool("list-build-options", false, "Print out a list of available build options for the iOS Xcode sub-project.")
	analyzeCmd.Flags().Bool("list-build-variants", false, "Print out a list of available build variants for the Android sub-project.")
	analyzeCmd.Flags().Bool("machine", false, "Outputs in a machine readable structured JSON format.")
	analyzeCmd.Flags().Bool("no-congratulate", false, "Show output even when there are no errors, warnings, hints, or lints. Ignored if \"--watch\" is specified.")
	analyzeCmd.Flags().Bool("no-current-package", false, "Analyze the current project, if applicable.")
	analyzeCmd.Flags().Bool("no-fatal-infos", false, "Treat info level issues as fatal.")
	analyzeCmd.Flags().Bool("no-fatal-warnings", false, "Treat warning level issues as fatal.")
	analyzeCmd.Flags().Bool("no-list-build-options", false, "Print out a list of available build options for the iOS Xcode sub-project.")
	analyzeCmd.Flags().Bool("no-plugins", false, "Whether to enable analyzer plugins.")
	analyzeCmd.Flags().Bool("no-preamble", false, "When analyzing the flutter repository, display the number of files that will be analyzed. Ignored if \"--watch\" is specified.")
	analyzeCmd.Flags().Bool("no-pub", false, "Whether to run \"flutter pub get\" before executing this command.")
	analyzeCmd.Flags().Bool("no-suggestions", false, "Show suggestions about the current flutter project.")
	analyzeCmd.Flags().Bool("output-app-link-settings", false, "Output a JSON with Android app link settings into a file. The \"--build-variant\" must also be set.")
	analyzeCmd.Flags().Bool("output-universal-link-settings", false, "Output a JSON with iOS Xcode universal link settings into a file. The \"--configuration\" and \"--target\" must be set.")
	analyzeCmd.Flags().Bool("plugins", false, "Whether to enable analyzer plugins.")
	analyzeCmd.Flags().Bool("preamble", false, "When analyzing the flutter repository, display the number of files that will be analyzed. Ignored if \"--watch\" is specified.")
	analyzeCmd.Flags().String("protocol-traffic-log", "", "The path to write the request and response protocol. This is only intended to be used for debugging the tooling.")
	analyzeCmd.Flags().Bool("pub", false, "Whether to run \"flutter pub get\" before executing this command.")
	analyzeCmd.Flags().Bool("suggestions", false, "Show suggestions about the current flutter project.")
	analyzeCmd.Flags().String("target", "", "Sets the iOS build target to be analyzed.")
	analyzeCmd.Flags().Bool("watch", false, "Run analysis continuously, watching the filesystem for changes.")
	analyzeCmd.Flags().String("write", "", "Also output the results to a file. This is useful with \"--watch\" if you want a file to always contain the latest results.")
	rootCmd.AddCommand(analyzeCmd)

	carapace.Gen(analyzeCmd).FlagCompletion(carapace.ActionMap{
		"write": carapace.ActionFiles(),
	})
}
