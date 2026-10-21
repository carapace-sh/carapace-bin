package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/flutter_completer/cmd/action"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new Flutter project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()

	createCmd.Flags().StringP("android-language", "a", "", "The language to use for Android-specific code, either Kotlin (recommended) or Java (legacy).")
	createCmd.Flags().String("description", "", "The description to use for your new Flutter project. This string ends up in the pubspec.yaml file.")
	createCmd.Flags().BoolP("empty", "e", false, "Specifies creating using an application template with a main.dart that is minimal, including no comments, as a starting point for a new application. Implies \"--template=app\".")
	createCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	createCmd.Flags().Bool("implementation-tests", false, "Include implementation tests that verify the template functions correctly. This is only intended to enable testing of the tool itself.")
	createCmd.Flags().String("initial-create-revision", "", "The Flutter SDK git commit hash to store in .migrate_config. This parameter is used by the tool internally and should generally not be used manually.")
	createCmd.Flags().StringP("ios-language", "i", "", "(deprecated) This option is deprecated and no longer has any effect. Swift is always used for iOS-specific code. This flag will be removed in a future version of Flutter.")
	createCmd.Flags().String("list-samples", "", "Specifies a JSON output file for a listing of Flutter code samples that can be created with \"--sample\".")
	createCmd.Flags().Bool("no-empty", false, "Specifies creating using an application template with a main.dart that is minimal, including no comments, as a starting point for a new application. Implies \"--template=app\".")
	createCmd.Flags().Bool("no-implementation-tests", false, "Include implementation tests that verify the template functions correctly. This is only intended to enable testing of the tool itself.")
	createCmd.Flags().Bool("no-offline", false, "When \"flutter pub get\" is run by the create command, this indicates whether to run it in offline mode or not. In offline mode, it will need to have all dependencies already available in the pub cache to succeed.")
	createCmd.Flags().Bool("no-overwrite", false, "When performing operations, overwrite existing files.")
	createCmd.Flags().Bool("no-pub", false, "Whether to run \"flutter pub get\" after the project has been created.")
	createCmd.Flags().Bool("no-skip-name-checks", false, "Allow the creation of applications and plugins with invalid names. This is only intended to enable testing of the tool itself.")
	createCmd.Flags().Bool("no-with-driver-test", false, "(deprecated) Historically, this added a flutter_driver dependency and generated a sample \"flutter drive\" test. Now it does nothing. Consider using the \"integration_test\" package: https://pub.dev/packages/integration_test")
	createCmd.Flags().Bool("offline", false, "When \"flutter pub get\" is run by the create command, this indicates whether to run it in offline mode or not. In offline mode, it will need to have all dependencies already available in the pub cache to succeed.")
	createCmd.Flags().String("org", "", "The organization responsible for your new Flutter project, in reverse domain name notation. This string is used in Java package names and as prefix in the iOS bundle identifier.")
	createCmd.Flags().Bool("overwrite", false, "When performing operations, overwrite existing files.")
	createCmd.Flags().StringSlice("platform", nil, "The platforms supported by this project. Platform folders (e.g. android/) will be generated in the target project. This argument only works when \"--template\" is set to app or plugin. When adding platforms to a plugin project, the pubspec.yaml will be updated with the requested platform. Adding desktop platforms requires the corresponding desktop config setting to be enabled.")
	createCmd.Flags().StringSlice("platforms", nil, "The platforms supported by this project. Platform folders (e.g. android/) will be generated in the target project. This argument only works when \"--template\" is set to app or plugin. When adding platforms to a plugin project, the pubspec.yaml will be updated with the requested platform. Adding desktop platforms requires the corresponding desktop config setting to be enabled.")
	createCmd.Flags().String("project-name", "", "The project name for this new Flutter project. This must be a valid dart package name.")
	createCmd.Flags().Bool("pub", false, "Whether to run \"flutter pub get\" after the project has been created.")
	createCmd.Flags().StringP("sample", "s", "", "Specifies the Flutter code sample to use as the \"main.dart\" for an application. Implies \"--template=app\". The value should be the sample ID of the desired sample from the API documentation website (https://api.flutter.dev/). An example can be found at: https://api.flutter.dev/flutter/widgets/SingleChildScrollView-class.html")
	createCmd.Flags().Bool("skip-name-checks", false, "Allow the creation of applications and plugins with invalid names. This is only intended to enable testing of the tool itself.")
	createCmd.Flags().StringP("template", "t", "", "Specify the type of project to create.")
	createCmd.Flags().Bool("with-driver-test", false, "(deprecated) Historically, this added a flutter_driver dependency and generated a sample \"flutter drive\" test. Now it does nothing. Consider using the \"integration_test\" package: https://pub.dev/packages/integration_test")
	rootCmd.AddCommand(createCmd)

	carapace.Gen(createCmd).FlagCompletion(carapace.ActionMap{
		"android-language": carapace.ActionValues("java", "kotlin"),
		"ios-language":     carapace.ActionValues("objc", "swift"),
		"platform":         carapace.ActionValues("ios", "android", "windows", "linux", "macos", "web", "darwin").UniqueList(","),
		"platforms":        carapace.ActionValues("ios", "android", "windows", "linux", "macos", "web", "darwin").UniqueList(","),
		"sample":           action.ActionSamples(),
		"template": carapace.ActionValuesDescribed(
			"app", "(default) Generate a Flutter application.",
			"module", "Generate a project to add a Flutter module to an existing Android or iOS application.",
			"package", "Generate a shareable Flutter project containing modular Dart code.",
			"package_ffi", "Generate a shareable Dart/Flutter project containing an API in Dart code with a platform-specific implementation through dart:ffi for Android, iOS, Linux, macOS, and Windows.",
			"plugin", "Generate a shareable Flutter project containing an API in Dart code with a platform-specific implementation through method channels for Android, iOS, Linux, macOS, Windows, web, or any combination of these.",
			"plugin_ffi", "(deprecated) Generate a shareable Flutter project containing an API in Dart code with a platform-specific implementation through dart:ffi for Android, iOS, Linux, macOS, Windows, or any combination of these. Use \"package_ffi\" instead.",
			"skeleton", "Formerly generated a list view / detail view Flutter application that followed some community best practices. For up to date resources, see https://flutter.github.io/samples, https://docs.flutter.dev/codelabs, and community resources such as https://flutter-builder.app/.",
		),
	})

	carapace.Gen(createCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
