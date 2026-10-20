package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var vars_promptCmd = &cobra.Command{
	Use:   "prompt",
	Short: "Ask for the prompt vars that have no saved answer, or set answers directly",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(vars_promptCmd).Standalone()

	vars_promptCmd.Flags().BoolP("help", "h", false, "Print help")
	varsCmd.AddCommand(vars_promptCmd)
}
