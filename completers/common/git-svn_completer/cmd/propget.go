package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var propgetCmd = &cobra.Command{
	Use:   "propget",
	Short: "Print the value of a property on a file or directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(propgetCmd).Standalone()

	propgetCmd.Flags().IntP("revision", "r", 0, "Refer to a specific revision")
	rootCmd.AddCommand(propgetCmd)

	carapace.Gen(propgetCmd).PositionalCompletion(
		carapace.ActionValues(),
		carapace.ActionFiles(),
	)
}
