package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Create executable and run immediately",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(runCmd).Standalone()

	runCmd.Flags().StringP("femit-bin", "", "", "Output machine code")
	runCmd.Flags().String("name", "", "Override root name")
	runCmd.Flags().Bool("strip", false, "Omit debug symbols")
	runCmd.Flags().StringP("target", "", "", "<arch><sub>-<os>-<abi>")
	runCmd.Flags().StringP("mcpu", "", "", "Specify target CPU and feature set")
	runCmd.Flags().StringP("O", "", "", "Choose what to optimize for")
	runCmd.Flags().BoolP("help", "h", false, "Print help")

	rootCmd.AddCommand(runCmd)

	carapace.Gen(runCmd).FlagCompletion(carapace.ActionMap{
		"O": carapace.ActionValues(
			"Debug",
			"ReleaseSafe",
			"ReleaseFast",
			"ReleaseSmall",
		),
	})

	carapace.Gen(runCmd).PositionalCompletion(
		carapace.ActionFiles(".zig"),
	)
}
