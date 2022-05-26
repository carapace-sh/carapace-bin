package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var store_dumpPathCmd = &cobra.Command{
	Use:   "dump-path",
	Short: "serialise a store path to stdout in NAR format",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(store_dumpPathCmd).Standalone()

	storeCmd.AddCommand(store_dumpPathCmd)

	addEvaluationFlags(store_dumpPathCmd)
	addFlakeFlags(store_dumpPathCmd)
	addLoggingFlags(store_dumpPathCmd)

	// TODO positional completion
}
