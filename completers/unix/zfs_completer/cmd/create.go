package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/zfs"
	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:     "create [-Pnpuv] [-o property=value]... filesystem|[-s] [-b blocksize] -V size volume",
	Short:   "create a new dataset",
	GroupID: "dataset",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(createCmd).Standalone()

	createCmd.Flags().BoolS("P", "P", false, "machine-parsable verbose output")
	createCmd.Flags().StringS("V", "V", "", "create a volume of the given size")
	createCmd.Flags().StringS("b", "b", "", "equivalent to -o volblocksize=blocksize")
	createCmd.Flags().BoolS("n", "n", false, "dry-run")
	createCmd.Flags().StringArrayS("o", "o", nil, "set property")
	createCmd.Flags().BoolS("p", "p", false, "create parent datasets")
	createCmd.Flags().BoolS("s", "s", false, "create sparse volume")
	createCmd.Flags().BoolS("u", "u", false, "do not mount the new filesystem")
	createCmd.Flags().BoolS("v", "v", false, "verbose")

	rootCmd.AddCommand(createCmd)

	carapace.Gen(createCmd).FlagCompletion(carapace.ActionMap{
		"b": carapace.ActionValues("512", "1K", "2K", "4K", "8K", "16K", "32K", "64K", "128K"),
		"o": zfs.ActionPropertyAssignments(),
	})
}
