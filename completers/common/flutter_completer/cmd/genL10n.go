package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var genL10nCmd = &cobra.Command{
	Use:   "gen-l10n",
	Short: "Generate localizations for the current project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(genL10nCmd).Standalone()

	genL10nCmd.Flags().String("arb-dir", "", "The directory where the template and translated arb files are located.")
	genL10nCmd.Flags().Bool("format", false, "When specified, the \"dart format\" command is run after generating the localization files.")
	genL10nCmd.Flags().String("gen-inputs-and-outputs-list", "", "When specified, the tool generates a JSON file containing the tool's inputs and outputs named gen_l10n_inputs_and_outputs.json. This can be useful for keeping track of which files of the Flutter project were used when generating the latest set of localizations. For example, the Flutter tool's build system uses this file to keep track of when to call gen_l10n during hot reload. The value of this option is the directory where the JSON file will be generated. When null, the JSON file will not be generated.")
	genL10nCmd.Flags().String("header", "", "The header to prepend to the generated Dart localizations files. This option takes in a string. For example, pass in \"/// All localized files.\" if you would like this string prepended to the generated Dart file. Alternatively, see the \"--header-file\" option to pass in a text file for longer headers.")
	genL10nCmd.Flags().String("header-file", "", "The header to prepend to the generated Dart localizations files. The value of this option is the name of the file that contains the header text which will be inserted at the top of each generated Dart file. Alternatively, see the \"--header\" option to pass in a string for a simpler header. This file should be placed in the directory specified in \"--arb-dir\".")
	genL10nCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	genL10nCmd.Flags().Bool("no-format", false, "When specified, the \"dart format\" command is run after generating the localization files.")
	genL10nCmd.Flags().Bool("no-nullable-getter", false, "Whether or not the localizations class getter is nullable. By default, this value is set to true so that Localizations.of(context) returns a nullable value for backwards compatibility. If this value is set to false, then a null check is performed on the returned value of Localizations.of(context), removing the need for null checking in user code.")
	genL10nCmd.Flags().Bool("no-relax-syntax", false, "When specified, the syntax will be relaxed so that the special character \"{\" is treated as a string if it is not followed by a valid placeholder and \"}\" is treated as a string if it does not close any previous \"{\" that is treated as a special character.")
	genL10nCmd.Flags().Bool("no-required-resource-attributes", false, "Requires all resource ids to contain a corresponding resource attribute. By default, simple messages will not require metadata, but it is highly recommended as this provides context for the meaning of a message to readers. Resource attributes are still required for plural messages.")
	genL10nCmd.Flags().Bool("no-suppress-warnings", false, "When specified, all warnings will be suppressed.")
	genL10nCmd.Flags().Bool("no-synthetic-package", false, "DEPRECATED. This flag cannot be enabled and should be removed.")
	genL10nCmd.Flags().Bool("no-use-deferred-loading", false, "Whether to generate the Dart localization file with locales imported as deferred, allowing for lazy loading of each locale in Flutter web. This can reduce a web app’s initial startup time by decreasing the size of the JavaScript bundle. When this flag is set to true, the messages for a particular locale are only downloaded and loaded by the Flutter app as they are needed. For projects with a lot of different locales and many localization strings, it can be an performance improvement to have deferred loading. For projects with a small number of locales, the difference is negligible, and might slow down the start up compared to bundling the localizations with the rest of the application. This flag does not affect other platforms such as mobile or desktop.")
	genL10nCmd.Flags().Bool("no-use-escaping", false, "Whether or not to use escaping for messages. By default, this value is set to false for backwards compatibility. Turning this flag on will cause the parser to treat any special characters contained within pairs of single quotes as normal strings and treat all consecutive pairs of single quotes as a single quote character.")
	genL10nCmd.Flags().Bool("no-use-named-parameters", false, "Whether or not to use named parameters for the generated localization methods.")
	genL10nCmd.Flags().Bool("nullable-getter", false, "Whether or not the localizations class getter is nullable. By default, this value is set to true so that Localizations.of(context) returns a nullable value for backwards compatibility. If this value is set to false, then a null check is performed on the returned value of Localizations.of(context), removing the need for null checking in user code.")
	genL10nCmd.Flags().String("output-class", "", "The Dart class name to use for the output localization and localizations delegate classes.")
	genL10nCmd.Flags().String("output-dir", "", "The directory where the generated localization classes will be written. The app must import the file specified in the \"--output-localization-file\" option from this directory. If unspecified, this defaults to the same directory as the input directory specified in \"--arb-dir\".")
	genL10nCmd.Flags().String("output-localization-file", "", "The filename for the output localization and localizations delegate classes.")
	genL10nCmd.Flags().StringSlice("preferred-supported-locales", nil, "The list of preferred supported locales for the application. By default, the tool will generate the supported locales list in alphabetical order. Use this flag if you would like to default to a different locale. For example, pass in \"en_US\" if you would like your app to default to American English on devices that support it. Pass this option multiple times to define multiple items.")
	genL10nCmd.Flags().String("project-dir", "", "When specified, the tool uses the path passed into this option as the directory of the root Flutter project. When null, the relative path to the present working directory will be used.")
	genL10nCmd.Flags().Bool("relax-syntax", false, "When specified, the syntax will be relaxed so that the special character \"{\" is treated as a string if it is not followed by a valid placeholder and \"}\" is treated as a string if it does not close any previous \"{\" that is treated as a special character.")
	genL10nCmd.Flags().Bool("required-resource-attributes", false, "Requires all resource ids to contain a corresponding resource attribute. By default, simple messages will not require metadata, but it is highly recommended as this provides context for the meaning of a message to readers. Resource attributes are still required for plural messages.")
	genL10nCmd.Flags().Bool("suppress-warnings", false, "When specified, all warnings will be suppressed.")
	genL10nCmd.Flags().Bool("synthetic-package", false, "DEPRECATED. This flag cannot be enabled and should be removed.")
	genL10nCmd.Flags().String("template-arb-file", "", "The template arb file that will be used as the basis for generating the Dart localization and messages files.")
	genL10nCmd.Flags().String("untranslated-messages-file", "", "The location of a file that describes the localization messages have not been translated yet. Using this option will create a JSON file at the target location, in the following format: \"locale\": [\"message_1\", \"message_2\" ... \"message_n\"] If this option is not specified, a summary of the messages that have not been translated will be printed on the command line.")
	genL10nCmd.Flags().Bool("use-deferred-loading", false, "Whether to generate the Dart localization file with locales imported as deferred, allowing for lazy loading of each locale in Flutter web. This can reduce a web app’s initial startup time by decreasing the size of the JavaScript bundle. When this flag is set to true, the messages for a particular locale are only downloaded and loaded by the Flutter app as they are needed. For projects with a lot of different locales and many localization strings, it can be an performance improvement to have deferred loading. For projects with a small number of locales, the difference is negligible, and might slow down the start up compared to bundling the localizations with the rest of the application. This flag does not affect other platforms such as mobile or desktop.")
	genL10nCmd.Flags().Bool("use-escaping", false, "Whether or not to use escaping for messages. By default, this value is set to false for backwards compatibility. Turning this flag on will cause the parser to treat any special characters contained within pairs of single quotes as normal strings and treat all consecutive pairs of single quotes as a single quote character.")
	genL10nCmd.Flags().Bool("use-named-parameters", false, "Whether or not to use named parameters for the generated localization methods.")
	rootCmd.AddCommand(genL10nCmd)
}
