package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var jsonschemaCmd = &cobra.Command{
	Use:     "jsonschema",
	Short:   "Outputs goreleaser's JSON schema",
	Aliases: []string{"schema"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(jsonschemaCmd).Standalone()

	jsonschemaCmd.Flags().StringP("output", "o", "", "Where to save the JSONSchema file")
	rootCmd.AddCommand(jsonschemaCmd)

	carapace.Gen(jsonschemaCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
	})
}
