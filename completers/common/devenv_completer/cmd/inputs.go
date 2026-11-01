package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inputsCmd = &cobra.Command{
	Use:   "inputs",
	Short: "Add an input to devenv.yaml",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inputsCmd).Standalone()

	rootCmd.AddCommand(inputsCmd)
}
