package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/apt_completer/cmd/common"
	"github.com/spf13/cobra"
)

var autopurgeCmd = &cobra.Command{
	Use:   "autopurge",
	Short: "automatically purge all unused packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autopurgeCmd).Standalone()

	common.AddGetFlags(autopurgeCmd)
	common.ActionInstallFlags(autopurgeCmd)
	rootCmd.AddCommand(autopurgeCmd)
}
