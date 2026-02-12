package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Print Atuin's shell init script",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().Bool("disable-ctrl-r", false, "Disable the binding of CTRL-R to atuin")
	initCmd.Flags().Bool("disable-up-arrow", false, "Disable the binding of the Up Arrow key to atuin")
	initCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).PositionalCompletion(
		carapace.ActionValues("zsh", "bash", "fish", "nu", "xonsh", "powershell"),
	)
}
