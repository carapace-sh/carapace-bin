package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var multiPackIndex_writeCmd = &cobra.Command{
	Use:   "write",
	Short: "Write a new MIDX file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(multiPackIndex_writeCmd).Standalone()

	multiPackIndex_writeCmd.Flags().Bool("bitmap", false, "write a multi-pack bitmap")
	multiPackIndex_writeCmd.Flags().Bool("no-bitmap", false, "do not write a multi-pack bitmap")
	multiPackIndex_writeCmd.Flags().String("preferred-pack", "", "specify the tie-breaking pack used when multiple packs contain the same object")
	multiPackIndex_writeCmd.Flags().String("refs-snapshot", "", "specify a file which contains a \"refs snapshot\" taken prior to repacking")
	multiPackIndex_writeCmd.Flags().Bool("stdin-packs", false, "provide pack index basenames over stdin")
	multiPackIndexCmd.AddCommand(multiPackIndex_writeCmd)

	carapace.Gen(multiPackIndex_writeCmd).FlagCompletion(carapace.ActionMap{
		"refs-snapshot": carapace.ActionFiles(),
	})
}
