package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var constraints_queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query the constraints fact database",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(constraints_queryCmd).Standalone()

	constraints_queryCmd.Flags().Bool("json", false, "Format the output as an NDJSON stream")
	constraintsCmd.AddCommand(constraints_queryCmd)
}
