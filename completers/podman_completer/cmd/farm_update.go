package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var farm_updateCmd = &cobra.Command{
	Use:   "update [options] FARM",
	Short: "Update an existing farm",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(farm_updateCmd).Standalone()

	farm_updateCmd.Flags().StringSliceP("add", "a", []string{}, "add system connection(s) to farm")
	farm_updateCmd.Flags().BoolP("default", "d", false, "set the given farm as the default farm")
	farm_updateCmd.Flags().StringSliceP("remove", "r", []string{}, "remove system connection(s) from farm")
	farmCmd.AddCommand(farm_updateCmd)
}
