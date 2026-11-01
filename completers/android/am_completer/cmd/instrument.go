package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var instrumentCmd = &cobra.Command{
	Use:   "instrument",
	Short: "Start an Instrumentation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(instrumentCmd).Standalone()

	instrumentCmd.Flags().String("abi", "", "launch the instrumented process with the selected `ABI`")
	instrumentCmd.Flags().StringP("argument", "e", "", "set argument `NAME` to `VALUE`")
	instrumentCmd.Flags().Bool("no-hidden-api-checks", false, "disable restrictions on use of hidden API")
	instrumentCmd.Flags().Bool("no-isolated-storage", false, "don't use isolated storage sandbox and mount full external storage")
	instrumentCmd.Flags().Bool("no-test-api-access", false, "do not allow access to test APIs, if hidden API checks are enabled")
	instrumentCmd.Flags().Bool("no-window-animation", false, "turn off window animations while running")
	instrumentCmd.Flags().StringP("profile", "p", "", "write profiling data to `FILE`")
	instrumentCmd.Flags().Bool("protobuf", false, "write output as protobuf to stdout (machine readable)")
	instrumentCmd.Flags().BoolP("raw-results", "r", false, "print raw results")
	instrumentCmd.Flags().String("user", "", "specify `USER_ID` instrumentation runs in")
	instrumentCmd.Flags().BoolP("wait", "w", false, "wait for instrumentation to finish before returning")

	rootCmd.AddCommand(instrumentCmd)

	carapace.Gen(instrumentCmd).FlagCompletion(carapace.ActionMap{
		"profile": carapace.ActionFiles(),
		"user": carapace.Batch(
			carapace.ActionValues("current"),
			android.ActionUsers(),
		).ToA(),
	})

	carapace.Gen(instrumentCmd).PositionalAnyCompletion(
		android.ActionComponents(),
	)
}
