package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/idf.py_completer/cmd/common"
	"github.com/spf13/cobra"
)

var efuseWriteProtectCmd = &cobra.Command{
	Use:   "efuse-write-protect",
	Short: "Write-protect eFuse",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efuseWriteProtectCmd).Standalone()
	common.AddEfuseFlags(efuseWriteProtectCmd)
	rootCmd.AddCommand(efuseWriteProtectCmd)
}
