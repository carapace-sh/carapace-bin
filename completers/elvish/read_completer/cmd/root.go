package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "read",
	Short: "Read a line from stdin",
	Long:  "https://elv.sh/ref/builtin.html#read",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().StringS("delimiter", "delimiter", "", "delimiter to use for read-upto")
	rootCmd.Flags().BoolS("line", "line", false, "read a line")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"delimiter": carapace.ActionValues().Usage("delimiter byte"),
	})
}
