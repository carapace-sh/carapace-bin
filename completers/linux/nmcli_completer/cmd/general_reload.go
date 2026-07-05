package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var general_reloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "Reload NetworkManager's configuration",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(general_reloadCmd).Standalone()

	generalCmd.AddCommand(general_reloadCmd)

	carapace.Gen(general_reloadCmd).PositionalCompletion(
		carapace.ActionValues("conf", "dns-rc", "dns-full").UniqueList(","),
	)
}
