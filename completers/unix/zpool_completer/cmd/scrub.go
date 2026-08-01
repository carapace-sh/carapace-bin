package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var scrubCmd = &cobra.Command{
	Use:     "scrub [-e|-p|-s|-C] [-w] [-S date] [-E date] -a|pool...",
	Short:   "start or manage a scrub",
	GroupID: "maintenance",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scrubCmd).Standalone()

	scrubCmd.Flags().BoolS("C", "C", false, "continue scrub from last saved txg")
	scrubCmd.Flags().StringS("E", "E", "", "end date for scrub range (YYYY-MM-DD HH:MM)")
	scrubCmd.Flags().StringS("S", "S", "", "start date for scrub range (YYYY-MM-DD HH:MM)")
	scrubCmd.Flags().BoolP("all", "a", false, "begin, pause, stop scrub on all pools")
	scrubCmd.Flags().BoolS("e", "e", false, "only scrub error blocks")
	scrubCmd.Flags().BoolS("p", "p", false, "pause scrub")
	scrubCmd.Flags().BoolS("s", "s", false, "stop scrub")
	scrubCmd.Flags().BoolS("w", "w", false, "wait until scrub completes")

	rootCmd.AddCommand(scrubCmd)

	carapace.Gen(scrubCmd).PositionalAnyCompletion(
		zfs.ActionPools(),
	)
}
