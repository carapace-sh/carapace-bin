package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/swift_completer/cmd/common"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Build and run an executable product",
	Args:  cobra.MinimumNArgs(0),
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()
	runCmd.Flags().SetInterspersed(false)

	common.AddPackageFlags(runCmd)

	runCmd.Flags().Bool("build-tests", false, "Build both source and test targets")
	runCmd.Flags().Bool("debugger", false, "Launch the executable in a debugger session")
	runCmd.Flags().BoolP("help", "h", false, "Show help information")
	runCmd.Flags().Bool("repl", false, "Launch Swift REPL for the package")
	runCmd.Flags().Bool("run", false, "Launch the executable with the provided arguments")
	runCmd.Flags().Bool("skip-build", false, "Skip building the executable product")
	runCmd.Flags().Bool("version", false, "Show the version")

	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return carapace.ActionExecCommand("swift", "package", "show-executables", "--format", "json")(func(output []byte) carapace.Action {
				return carapace.ActionValues()
			})
		}),
	)

	carapace.Gen(runCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
