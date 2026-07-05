package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var apfsSetPassphraseHintCmd = &cobra.Command{
	Use:   "setPassphraseHint",
	Short: "Set or clear a passphrase hint",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apfsSetPassphraseHintCmd).Standalone()
	apfsCmd.AddCommand(apfsSetPassphraseHintCmd)
	carapace.Gen(apfsSetPassphraseHintCmd).PositionalAnyCompletion(fs.ActionBlockDevices())
}
