package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hazmat_n1Cmd = &cobra.Command{
	Use:   "n1 [cmd...]",
	Short: "run the command once per filename",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hazmat_n1Cmd).Standalone()

	carapace.Gen(hazmat_n1Cmd).PositionalAnyCompletion(
		carapace.ActionExecutables(),
	)

	hazmatCmd.AddCommand(hazmat_n1Cmd)
}
