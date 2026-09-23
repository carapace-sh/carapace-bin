package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/android"
	"github.com/spf13/cobra"
)

var listInstrumentationCmd = &cobra.Command{
	Use:   "instrumentation",
	Short: "Print all test packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listInstrumentationCmd).Standalone()

	listInstrumentationCmd.Flags().BoolP("file", "f", false, "dump the name of the .apk file containing the test package")

	listCmd.AddCommand(listInstrumentationCmd)

	carapace.Gen(listInstrumentationCmd).PositionalCompletion(
		android.ActionPackages(),
	)
}
