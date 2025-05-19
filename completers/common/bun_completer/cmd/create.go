package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new project from a template",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()

	rootCmd.AddCommand(createCmd)

	carapace.Gen(createCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
