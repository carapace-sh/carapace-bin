package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hexInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Print shell code to initialize atuin-hex on shell startup",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hexInitCmd).Standalone()

	hexInitCmd.Flags().BoolP("help", "h", false, "Print help")
	hexCmd.AddCommand(hexInitCmd)

	carapace.Gen(hexInitCmd).PositionalCompletion(
		carapace.ActionValues("zsh", "bash", "fish", "nu"),
	)
}
