package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vars_unsetCmd = &cobra.Command{
	Use:     "unset",
	Short:   "Forget saved answers so they are asked for again",
	Aliases: []string{"rm", "remove"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vars_unsetCmd).Standalone()

	vars_unsetCmd.Flags().BoolP("help", "h", false, "Print help")
	varsCmd.AddCommand(vars_unsetCmd)
}
