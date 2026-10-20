package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bootstrap_miseShellActivateCmd = &cobra.Command{
	Use:     "mise-shell-activate",
	Short:   "Manage mise shell activation from `[bootstrap.mise_shell_activate]`",
	Aliases: []string{"shell"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootstrap_miseShellActivateCmd).Standalone()

	bootstrap_miseShellActivateCmd.Flags().BoolP("help", "h", false, "Print help")
	bootstrapCmd.AddCommand(bootstrap_miseShellActivateCmd)
}
