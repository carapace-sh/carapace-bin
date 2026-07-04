package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cd",
	Short: "Change the current working directory",
	Long:  "https://zsh.sourceforge.io/Doc/Release/Shell-Builtin-Commands.html#index-cd",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.Flags().BoolS("L", "L", false, "force symbolic links to be followed")
	rootCmd.Flags().BoolS("P", "P", false, "use the physical directory structure without following symbolic links")
	rootCmd.Flags().BoolS("e", "e", false, "if the -P option is supplied, exit with a non-zero status if the current working directory cannot be determined successfully")
	rootCmd.Flags().BoolS("q", "q", false, "suppress hook function chpwd and chpwd_functions")
	rootCmd.Flags().BoolS("s", "s", false, "refuse to change the current directory if the target directory is not found in cdpath")

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
