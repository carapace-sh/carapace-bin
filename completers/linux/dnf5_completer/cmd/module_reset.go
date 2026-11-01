package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var moduleResetCmd = &cobra.Command{
	Use:   "reset [options] <module-spec>...",
	Short: "reset module state so it's no longer enabled or disabled",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(moduleResetCmd).Standalone()

	moduleResetCmd.Flags().Bool("skip-unavailable", false, "Allow skipping unavailable packages")

	moduleCmd.AddCommand(moduleResetCmd)
}
