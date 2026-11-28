package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/idf.py_completer/cmd/common"
	"github.com/spf13/cobra"
)

var efuseReadProtectCmd = &cobra.Command{
	Use:   "efuse-read-protect",
	Short: "Read-protect eFuse",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efuseReadProtectCmd).Standalone()

	common.AddEfuseFlags(efuseReadProtectCmd)
	rootCmd.AddCommand(efuseReadProtectCmd)
}
