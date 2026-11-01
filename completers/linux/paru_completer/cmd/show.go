package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/paru_completer/cmd/common"
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

	showCmd.Flags().BoolP("complete", "c", false, "Used for completions")
	showCmd.Flags().CountP("news", "w", "Print arch news")
	showCmd.Flags().BoolP("order", "o", false, "Print buildorder for targets")
	showCmd.Flags().BoolP("stats", "s", false, "Display system package statistics")
	common.AddNewFlags(showCmd)
}
