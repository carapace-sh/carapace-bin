package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var multiPackIndex_compactCmd = &cobra.Command{
	Use:   "compact",
	Short: "Compact a multi-pack-index",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(multiPackIndex_compactCmd).Standalone()

	multiPackIndex_compactCmd.Flags().String("base", "", "base MIDX for incremental writes")
	multiPackIndex_compactCmd.Flags().Bool("bitmap", false, "write multi-pack bitmap")
	multiPackIndex_compactCmd.Flags().Bool("incremental", false, "write a new incremental MIDX")
	multiPackIndex_compactCmd.Flags().Bool("no-write-chain-file", false, "do not write the multi-pack-index chain file")
	multiPackIndex_compactCmd.Flags().String("object-dir", "", "use given directory for the location of Git objects")
	multiPackIndex_compactCmd.Flags().Bool("progress", false, "force progress reporting")
	multiPackIndexCmd.AddCommand(multiPackIndex_compactCmd)

	carapace.Gen(multiPackIndex_compactCmd).FlagCompletion(carapace.ActionMap{
		"object-dir": carapace.ActionDirectories(),
	})
}
