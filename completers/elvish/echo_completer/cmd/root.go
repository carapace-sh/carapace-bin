package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "echo",
	Short: "Print arguments with trailing newline",
	Long:  "https://elv.sh/ref/builtin.html#echo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("sep", "sep", " ", "separator to use between arguments")

	carapace.Gen(rootCmd).PositionalAnyCompletion(
		carapace.ActionValues().Usage("value to print"),
	)
}
