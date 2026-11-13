package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hazmat_ignoreExitCodeCmd = &cobra.Command{
	Use:   "ignore-exit-code [cmd...]",
	Short: "run the command but ignore the exit code",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hazmat_ignoreExitCodeCmd).Standalone()

	carapace.Gen(hazmat_ignoreExitCodeCmd).PositionalAnyCompletion(
		carapace.ActionExecutables(),
	)

	hazmatCmd.AddCommand(hazmat_ignoreExitCodeCmd)
}
