package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var secret_generateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generate a random secret value",
	Aliases: []string{"gen"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secret_generateCmd).Standalone()

	secret_generateCmd.Flags().Int("length", 32, "Number of random bytes")
	secret_generateCmd.Flags().StringP("output", "o", "", "Write to file instead of stdout")
	secretCmd.AddCommand(secret_generateCmd)

	carapace.Gen(secret_generateCmd).FlagCompletion(carapace.ActionMap{
		"output": carapace.ActionFiles(),
	})
}
