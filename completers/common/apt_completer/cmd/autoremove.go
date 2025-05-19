package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/apt_completer/cmd/common"
	"github.com/spf13/cobra"
)

var autoremoveCmd = &cobra.Command{
	Use:   "autoremove",
	Short: "automatically remove all unused packages",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(autoremoveCmd).Standalone()

	common.AddGetFlags(autoremoveCmd)
	common.ActionInstallFlags(autoremoveCmd)
	rootCmd.AddCommand(autoremoveCmd)
}
