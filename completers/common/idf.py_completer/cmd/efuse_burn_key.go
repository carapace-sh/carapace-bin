package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/idf.py_completer/cmd/common"
	"github.com/spf13/cobra"
)

var efuseBurnKeyCmd = &cobra.Command{
	Use:   "efuse-burn-key",
	Short: "Burn eFuse key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(efuseBurnKeyCmd).Standalone()

	common.AddEfuseFlags(efuseBurnKeyCmd)
	efuseBurnKeyCmd.Flags().Bool("force-write-always", false, "Write the eFuse even if already written or write protected")
	efuseBurnKeyCmd.Flags().Bool("no-protect-key", false, "Disable default read- and write-protecting of the key")
	efuseBurnKeyCmd.Flags().Bool("show-sensitive-info", false, "Show data to be burned")
	rootCmd.AddCommand(efuseBurnKeyCmd)
}
