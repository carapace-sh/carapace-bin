package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/idf.py_completer/cmd/common"
	"github.com/spf13/cobra"
)

var encryptedFlashCmd = &cobra.Command{
	Use:   "encrypted-flash",
	Short: "Flash the encrypted project",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(encryptedFlashCmd).Standalone()

	common.AddFlashFlags(encryptedFlashCmd)
	rootCmd.AddCommand(encryptedFlashCmd)
}
