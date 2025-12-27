package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/linux/apt_completer/cmd/common"
	"github.com/spf13/cobra"
)

var purgeCmd = &cobra.Command{
	Use:   "purge",
	Short: "remove package including config and data",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(purgeCmd).Standalone()

	common.AddGetFlags(purgeCmd)
	common.ActionInstallFlags(purgeCmd)
	rootCmd.AddCommand(purgeCmd)
}
