package common

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func AddRandomStateFlags(cmd *cobra.Command) {
	cmd.Flags().StringS("rand", "rand", "", "Load the given file(s) into the random number generator")
	cmd.Flags().StringS("writerand", "writerand", "", "Write random data to the specified file")

	carapace.Gen(cmd).FlagCompletion(carapace.ActionMap{
		"rand":      carapace.ActionFiles(),
		"writerand": carapace.ActionFiles(),
	})
}
