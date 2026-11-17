package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serverUninstallCmd = &cobra.Command{
	Use:   "uninstall [ssh://[USER@]HOST[:PORT]]",
	Short: "Stop starting your server on its own",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serverUninstallCmd).Standalone()

	serverUninstallCmd.Flags().Bool("json", false, "output what the command did as JSON")
	serverUninstallCmd.Flags().Bool("user", false, "undo a rex server install --user")
	serverUninstallCmd.Flags().Bool("yes", false, "answer yes to every question")

	serverCmd.AddCommand(serverUninstallCmd)
}
