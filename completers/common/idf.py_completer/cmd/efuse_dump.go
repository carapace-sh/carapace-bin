package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/idf.py_completer/cmd/common"
	"github.com/spf13/cobra"
)

var efuseDumpCmd = &cobra.Command{
	Use:   "efuse-dump",
	Short: "Dump eFuse blocks",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efuseDumpCmd).Standalone()

	common.AddEfuseFlags(efuseDumpCmd)
	efuseDumpCmd.Flags().String("file-name", "", "Saves dump for each block into separate file")
	rootCmd.AddCommand(efuseDumpCmd)

	carapace.Gen(efuseDumpCmd).FlagCompletion(carapace.ActionMap{
		"file-name": carapace.ActionFiles(),
	})
}
