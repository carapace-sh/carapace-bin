package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var template_lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Parse Benthos templates and report any linting errors",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(template_lintCmd).Standalone()

	templateCmd.AddCommand(template_lintCmd)

	carapace.Gen(template_lintCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
