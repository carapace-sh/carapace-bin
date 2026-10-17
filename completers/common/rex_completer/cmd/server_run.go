package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/common"
	"github.com/spf13/cobra"
)

var serverRunCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a foreground server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverRunCmd).Standalone()
	common.AddServerFlags(serverRunCmd)

	serverCmd.AddCommand(serverRunCmd)
}
