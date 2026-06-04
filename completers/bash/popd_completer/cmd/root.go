package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "popd",
	Short: "Remove directories from directory stack",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Directory-Stack-Builtins.html#index-popd",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("n", "n", false, "suppress the normal change of directory when removing directories from the stack")
}
