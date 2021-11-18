package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Check whether the configuration is valid",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(validateCmd).Standalone()

	validateCmd.Flags().Bool("json", false, "Produce output in a machine-readable JSON format.")
	validateCmd.Flags().Bool("no-color", false, "If specified, output won't contain any color.")
	rootCmd.AddCommand(validateCmd)
}
