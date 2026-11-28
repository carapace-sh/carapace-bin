package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/idf.py_completer/cmd/common"
	"github.com/spf13/cobra"
)

var encryptedAppFlashCmd = &cobra.Command{
	Use:   "encrypted-app-flash",
	Short: "Flash the encrypted app only",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(encryptedAppFlashCmd).Standalone()

	common.AddFlashFlags(encryptedAppFlashCmd)
	rootCmd.AddCommand(encryptedAppFlashCmd)
}
