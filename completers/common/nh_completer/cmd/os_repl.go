package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var os_replCmd = &cobra.Command{
	Use:   "repl",
	Short: "Load system in a repl",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(os_replCmd).Standalone()

	os_replCmd.Flags().StringP("hostname", "H", "", "Select this hostname from nixosConfigurations")

	osCmd.AddCommand(os_replCmd)
}
