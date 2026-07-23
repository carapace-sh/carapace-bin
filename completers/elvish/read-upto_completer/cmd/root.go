package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "read-upto",
	Short: "Read bytes from stdin until a terminator",
	Long:  "https://elv.sh/ref/builtin.html#read-upto",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValues().Usage("terminator byte"),
	)
}
