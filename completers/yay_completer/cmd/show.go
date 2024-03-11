package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var showCmd = &cobra.Command{
	Use:     "show",
	Aliases: []string{"P"},
	Short:   "Print information",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(showCmd).Standalone()

	showCmd.Flags().BoolP("-c", "--complete", false, "Used for completions")
	showCmd.Flags().BoolP("-d", "--defaultconfig", false, "Print default yay configuration")
	showCmd.Flags().BoolP("-g", "--currentconfig", false, "Print current yay configuration")
	showCmd.Flags().BoolP("-s", "--stats", false, "Display system package statistics")
	showCmd.Flags().BoolP("-w", "--news", false, "Print arch news")
}
