package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/idf.py_completer/cmd/common"
	"github.com/spf13/cobra"
)

var efuseBurnCmd = &cobra.Command{
	Use:   "efuse-burn",
	Short: "Burn eFuse",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efuseBurnCmd).Standalone()

	common.AddEfuseFlags(efuseBurnCmd)
	rootCmd.AddCommand(efuseBurnCmd)
}
