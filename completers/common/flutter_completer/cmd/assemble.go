package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/flutter_completer/cmd/action"
	"github.com/spf13/cobra"
)

var assembleCmd = &cobra.Command{
	Use:   "assemble",
	Short: "Assemble and build Flutter resources",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(assembleCmd).Standalone()

	assembleCmd.Flags().String("build-inputs", "", "A file path where a newline-separated file containing all inputs used will be written after a build. This file is not included as a build input or output. This file is not written if the build fails for any reason.")
	assembleCmd.Flags().String("build-outputs", "", "A file path where a newline-separated file containing all outputs created will be written after a build. This file is not included as a build input or output. This file is not written if the build fails for any reason.")
	assembleCmd.Flags().StringArrayP("dart-define", "D", nil, "Additional key-value pairs that will be available as constants from the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple defines can be passed by repeating \"--dart-define\".")
	assembleCmd.Flags().StringSlice("dart-define-from-file", nil, "The path of a .json or .env file containing key-value pairs that will be available as environment variables. These can be accessed using the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple define files can be passed by repeating \"--dart-define-from-file\".")
	assembleCmd.Flags().StringArray("dart-defines", nil, "Additional key-value pairs that will be available as constants from the String.fromEnvironment, bool.fromEnvironment, and int.fromEnvironment constructors. Multiple defines can be passed by repeating \"--dart-define\".")
	assembleCmd.Flags().StringSliceP("define", "d", nil, "DEPRECATED. Use \"--dart-define\" or \"-D\" instead for consistency. Allows passing configuration to a target, as in \"--define=target=key=value\".")
	assembleCmd.Flags().String("depfile", "", "A file path where a depfile will be written. This contains all build inputs and outputs in a Make-style syntax.")
	assembleCmd.Flags().StringSlice("extra-front-end-options", nil, "A comma-separated list of additional command line arguments that will be passed directly to the Dart front end. For example, \"--extra-front-end-options=--enable-experiment=nonfunction-type-aliases\".")
	assembleCmd.Flags().StringSlice("extra-gen-snapshot-options", nil, "A comma-separated list of additional command line arguments that will be passed directly to the Dart native compiler. (Requires the \"--release\", \"--profile\", or \"--jit-release\" flag.)")
	assembleCmd.Flags().BoolP("help", "h", false, "Print this usage information.")
	assembleCmd.Flags().StringSliceP("input", "i", nil, "Allows passing additional inputs with \"--input=key=value\". Unlike defines, additional inputs do not generate a new configuration; instead they are treated as dependencies of the targets that use them.")
	assembleCmd.Flags().StringP("output", "o", "", "A directory where output files will be written. Must be either absolute or relative from the root of the current Flutter project.")
	assembleCmd.Flags().String("performance-measurement-file", "", "Output individual target performance to a JSON file.")
	assembleCmd.Flags().String("resource-pool-size", "", "The maximum number of concurrent tasks the build system will run.")
	rootCmd.AddCommand(assembleCmd)

	carapace.Gen(assembleCmd).FlagCompletion(carapace.ActionMap{
		"build-inputs":                 carapace.ActionFiles(),
		"build-outputs":                carapace.ActionFiles(),
		"depfile":                      carapace.ActionFiles(),
		"output":                       carapace.ActionDirectories(),
		"performance-measurement-file": carapace.ActionFiles(),
	})

	carapace.Gen(assembleCmd).PositionalCompletion(
		action.ActionTargets(),
	)
}
