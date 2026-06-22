package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/rails_completer/cmd/common"
	"github.com/spf13/cobra"
)

var standaloneEnvCmd = &cobra.Command{
	Use:   "standalone_env",
	Short: "Standalone Rails commands that accept --environment (boot, initializers)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(standaloneEnvCmd).Standalone()

	standaloneEnvCmd.Flags().BoolP("help", "h", false, "Show help")

	common.AddEnvironmentFlag(standaloneEnvCmd, "development")
}
