package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:     "get [-Hp] [-o field,...] all|property[,...] [pool...]",
	Short:   "get pool properties",
	GroupID: "property",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(getCmd).Standalone()

	getCmd.Flags().BoolS("H", "H", false, "scripting mode")
	getCmd.Flags().BoolP("json", "j", false, "JSON output")
	getCmd.Flags().Bool("json-int", false, "display numbers in integer format in JSON output")
	getCmd.Flags().Bool("json-pool-key-guid", false, "use pool GUID as key for pool objects in JSON output")
	getCmd.Flags().StringS("o", "o", "", "columns to display")
	getCmd.Flags().BoolS("p", "p", false, "display exact values")

	rootCmd.AddCommand(getCmd)

	carapace.Gen(getCmd).FlagCompletion(carapace.ActionMap{
		"o": carapace.ActionValues("name", "property", "value", "source").UniqueList(","),
	})

	carapace.Gen(getCmd).PositionalCompletion(
		carapace.ActionMultiParts(",", func(c carapace.Context) carapace.Action {
			return carapace.Batch(
				carapace.ActionValues("all"),
				zfs.ActionPoolProperties(),
				zfs.ActionReadonlyPoolProperties(),
			).ToA().FilterArgs()
		}),
		zfs.ActionPools(),
	)

	carapace.Gen(getCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) > 1 {
				return carapace.Batch(
					carapace.ActionValues("all-vdevs"),
					zfs.ActionPoolDevices(c.Args[1]),
				).ToA()
			}
			return carapace.ActionValues()
		}),
	)
}
