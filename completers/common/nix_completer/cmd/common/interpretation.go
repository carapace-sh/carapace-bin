package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

// AddInterpretationFlags adds flags for interpreting installables
// as expressions or files (SourceExprCommand in nix source).
func AddInterpretationFlags(cmd *cobra.Command) {
	cmd.Flags().String("expr", "", "Interpret installables as attribute paths relative to the Nix expression expr")
	cmd.Flags().StringP("file", "f", "", "Interpret installables as attribute paths relative to the Nix expression stored in file")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
	})
}
