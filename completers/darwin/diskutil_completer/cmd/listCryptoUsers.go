package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var apfsListCryptoUsersCmd = &cobra.Command{
	Use:   "listCryptoUsers",
	Short: "List cryptographic users on a volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apfsListCryptoUsersCmd).Standalone()
	apfsCmd.AddCommand(apfsListCryptoUsersCmd)
	carapace.Gen(apfsListCryptoUsersCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
