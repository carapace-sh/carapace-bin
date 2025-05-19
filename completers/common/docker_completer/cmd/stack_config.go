package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var stack_configCmd = &cobra.Command{
	Use:   "config [OPTIONS]",
	Short: "Outputs the final config file, after doing merges and interpolations",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(stack_configCmd).Standalone()

	stack_configCmd.Flags().StringSliceP("compose-file", "c", nil, "Path to a Compose file, or \"-\" to read from stdin")
	stack_configCmd.Flags().Bool("skip-interpolation", false, "Skip interpolation and output only merged config")
	stackCmd.AddCommand(stack_configCmd)
}
