package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var show_profileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Show the default profile used for the `rustup install` command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(show_profileCmd).Standalone()

	show_profileCmd.Flags().BoolP("help", "h", false, "Print help")
	showCmd.AddCommand(show_profileCmd)
}
