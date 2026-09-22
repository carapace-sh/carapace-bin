package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keys_syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Sync program `declare_id!` pubkeys with the program's actual pubkey",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keys_syncCmd).Standalone()

	keys_syncCmd.Flags().BoolP("help", "h", false, "Print help")
	keys_syncCmd.Flags().StringP("program-name", "p", "", "Only sync the given program instead of all programs")
	keysCmd.AddCommand(keys_syncCmd)
}
