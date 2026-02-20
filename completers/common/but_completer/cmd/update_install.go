package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var update_installCmd = &cobra.Command{
	Use:   "install",
	Short: "Install or update the GitButler desktop application",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(update_installCmd).Standalone()

	update_installCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	updateCmd.AddCommand(update_installCmd)
}
