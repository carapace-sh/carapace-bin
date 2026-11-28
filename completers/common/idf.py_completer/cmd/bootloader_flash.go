package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/idf.py_completer/cmd/common"
	"github.com/spf13/cobra"
)

var bootloaderFlashCmd = &cobra.Command{
	Use:   "bootloader-flash",
	Short: "Flash bootloader only",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bootloaderFlashCmd).Standalone()
	common.AddFlashFlags(bootloaderFlashCmd)
	rootCmd.AddCommand(bootloaderFlashCmd)
}
