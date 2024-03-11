package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/supervisor"
	"github.com/spf13/cobra"
)

var tailCmd = &cobra.Command{
	Use:   "tail",
	Short: "Output the last part of process log",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(tailCmd).Standalone()
	tailCmd.Flags().BoolS("f", "f", false, "continuous tail")

	rootCmd.AddCommand(tailCmd)

	carapace.Gen(tailCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			return supervisor.ActionProcesses(rootCmd.Flag("configuration").Value.String())
		}),
		carapace.ActionValues("stdout", "stderr"),
	)
}
