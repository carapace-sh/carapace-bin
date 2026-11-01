package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var propsetCmd = &cobra.Command{
	Use:   "propset",
	Short: "Set the value of a property on a file or directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(propsetCmd).Standalone()
	rootCmd.AddCommand(propsetCmd)

	carapace.Gen(propsetCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionValues(),
		carapace.ActionFiles(),
	)
}
