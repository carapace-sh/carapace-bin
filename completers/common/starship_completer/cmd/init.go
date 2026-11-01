package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Prints the shell function used to execute starship",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().BoolP("help", "h", false, "Print help")
	initCmd.Flags().Bool("print-full-init", false, "Print the main initialization script (as opposed to the init stub)")
	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).PositionalCompletion(
		carapace.ActionValues("bash", "cmd", "elvish", "fish", "ion", "nu", "powershell", "tcsh", "xonsh", "zsh"),
	)
}
