package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var apps_spec_validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate an application spec",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apps_spec_validateCmd).Standalone()
	apps_spec_validateCmd.Flags().Bool("schema-only", false, "Only validate the spec schema and not the correctness of the spec.")
	apps_specCmd.AddCommand(apps_spec_validateCmd)

	// TODO positional
}
