package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pushd",
	Short: "Add directories to the top of the directory stack",
	Long:  "https://www.gnu.org/software/bash/manual/html_node/Directory-Stack-Builtins.html#index-pushd",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("n", "n", false, "suppress the normal change of directory when adding directories to the stack")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionDirectories().FilterArgs(),
	)
}
