package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rex_completer/cmd/common"
	"github.com/spf13/cobra"
)

var serverStartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start a persistent server",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverStartCmd).Standalone()
	common.AddServerFlags(serverStartCmd)
	serverStartCmd.Flags().Bool("json", false, "output server information as JSON")
	serverStartCmd.Flags().String("log-max-size", "", "limit the active detached log to SIZE by clearing old contents")

	serverCmd.AddCommand(serverStartCmd)
}
