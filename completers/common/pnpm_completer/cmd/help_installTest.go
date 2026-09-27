package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_installTestCmd = &cobra.Command{
	Use:   "install-test",
	Short: "Runs a `pnpm install` followed immediately by a `pnpm test`. Accepts the same arguments as `pnpm install`, plus `--no-bail` to continue running workspace tests after a failure",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_installTestCmd).Standalone()

	helpCmd.AddCommand(help_installTestCmd)
}
