package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var farm_removeCmd = &cobra.Command{
	Use:     "remove [options] [FARM...]",
	Short:   "Remove one or more farms",
	Aliases: []string{"rm"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(farm_removeCmd).Standalone()

	farm_removeCmd.Flags().BoolP("all", "a", false, "Remove all farms")
	farmCmd.AddCommand(farm_removeCmd)
}
