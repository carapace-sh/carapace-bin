package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var pairableCmd = &cobra.Command{
	Use:   "pairable [on|off]",
	Short: "Set controller pairable mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pairableCmd).Standalone()
	rootCmd.AddCommand(pairableCmd)

	carapace.Gen(pairableCmd).PositionalCompletion(
		carapace.ActionValues("on", "off").StyleF(style.ForKeyword),
	)
}
