package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/anchor"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:     "test",
	Short:   "Runs integration tests",
	Aliases: []string{"t"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(testCmd).Standalone()

	testCmd.Flags().Bool("detach", false, "Flag to keep the local validator running after tests to be able to check the transactions")
	testCmd.Flags().StringSliceP("env", "e", nil, "Environment variables to pass into the docker container")
	testCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	testCmd.Flags().Bool("no-idl", false, "Do not build the IDL")
	testCmd.Flags().Bool("profile", false, "Profile each test: record per-test SBF register traces and render flamegraph SVGs under target/anchor-v2-profile")
	testCmd.Flags().StringP("program-name", "p", "", "Build and test only this program")
	testCmd.Flags().StringSlice("run", nil, "Run the test suites under the specified path")
	testCmd.Flags().String("script", "", "Name of the script to run from [scripts] section (defaults to \"test\")")
	testCmd.Flags().Bool("skip-build", false, "Flag to skip building the program in the workspace, use this to save time when running test and the program code is not altered")
	testCmd.Flags().Bool("skip-deploy", false, "Use this flag if you want to run tests against previously deployed programs")
	testCmd.Flags().Bool("skip-lint", false, "True if the build should not fail even if there are no \"CHECK\" comments where normally required")
	testCmd.Flags().Bool("skip-local-validator", false, "Flag to skip starting a local validator, if the configured cluster url is a localnet")
	testCmd.Flags().String("validator", "surfpool", "Validator type to use for local testing")
	rootCmd.AddCommand(testCmd)

	carapace.Gen(testCmd).FlagCompletion(carapace.ActionMap{
		"program-name": anchor.ActionPrograms(),
		"script":       anchor.ActionScripts(),
		"validator": carapace.ActionValuesDescribed(
			"surfpool", "Use Surfpool validator (default)",
			"legacy", "Use Solana test validator",
		),
	})
}
