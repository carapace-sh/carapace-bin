package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var codegenCmd = &cobra.Command{
	Use:   "--codegen [spec]",
	Short: "",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		return codegen(args[0])
	},
}

func init() {
	carapace.Gen(codegenCmd).Standalone()

	carapace.Gen(codegenCmd).PositionalCompletion(
		carapace.ActionFiles(".yaml"),
	)
}
