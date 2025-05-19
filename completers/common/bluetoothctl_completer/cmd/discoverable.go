package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var discoverableCmd = &cobra.Command{
	Use:   "discoverable [on|off]",
	Short: "Set controller discoverable mode",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discoverableCmd).Standalone()
	rootCmd.AddCommand(discoverableCmd)

	carapace.Gen(discoverableCmd).PositionalCompletion(
		carapace.ActionValues("on", "off").StyleF(style.ForKeyword),
	)
}
