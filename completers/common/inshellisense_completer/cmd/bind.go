package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bindCmd = &cobra.Command{
	Use:   "bind",
	Short: "adds keybindings to the selected shell",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bindCmd).Standalone()

	bindCmd.Flags().BoolP("help", "h", false, "display help for command")
	rootCmd.AddCommand(bindCmd)
}
