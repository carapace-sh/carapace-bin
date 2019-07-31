package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var programCmd = &cobra.Command{
	Use:     "program [-jn] [-t instruction-limit] [-m memlimit] pool script [arg...]",
	Short:   "execute a ZFS channel program",
	GroupID: "misc",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(programCmd).Standalone()

	programCmd.Flags().BoolP("json", "j", false, "display channel program output in JSON format")
	programCmd.Flags().StringS("m", "m", "", "memory limit in bytes")
	programCmd.Flags().BoolS("n", "n", false, "dry-run")
	programCmd.Flags().StringS("t", "t", "", "limit the number of Lua instructions to execute")

	rootCmd.AddCommand(programCmd)

	carapace.Gen(programCmd).PositionalCompletion(
		zfs.ActionPools(),
		carapace.ActionFiles(".lua"),
	)
}
