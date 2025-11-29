package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var genDocsCmd = &cobra.Command{
	Use:    "gen-docs <path>",
	Short:  "[Internal] Generate documentation for the CLI",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(genDocsCmd).Standalone()

	rootCmd.AddCommand(genDocsCmd)

	carapace.Gen(genDocsCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
