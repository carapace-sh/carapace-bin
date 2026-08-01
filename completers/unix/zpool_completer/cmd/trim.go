package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var trimCmd = &cobra.Command{
	Use:     "trim [-dw] [-r rate] [-c|-s] -a|pool [device...]",
	Short:   "initiate or manage TRIM on devices",
	GroupID: "maintenance",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(trimCmd).Standalone()

	trimCmd.Flags().BoolP("all", "a", false, "perform TRIM operation on all pools")
	trimCmd.Flags().BoolP("cancel", "c", false, "cancel TRIM")
	trimCmd.Flags().StringP("rate", "r", "", "TRIM rate limit")
	trimCmd.Flags().BoolP("secure", "d", false, "secure TRIM")
	trimCmd.Flags().BoolP("suspend", "s", false, "suspend TRIM")
	trimCmd.Flags().BoolP("wait", "w", false, "wait until TRIM completes")

	rootCmd.AddCommand(trimCmd)

	carapace.Gen(trimCmd).PositionalCompletion(
		zfs.ActionPools(),
	)

	carapace.Gen(trimCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 0 {
				return zfs.ActionPoolDevices(c.Args[0])
			}
			return carapace.ActionValues()
		}),
	)
}
