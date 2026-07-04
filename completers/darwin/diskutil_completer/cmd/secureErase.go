package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/fs"
	"github.com/spf13/cobra"
)

var secureEraseCmd = &cobra.Command{
	Use:   "secureErase",
	Short: "Securely erase a disk or volume",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(secureEraseCmd).Standalone()
	rootCmd.AddCommand(secureEraseCmd)
	carapace.Gen(secureEraseCmd).PositionalCompletion(
		carapace.ActionValues("0", "1", "2", "3", "4"),
		fs.ActionBlockDevices(),
	)
}
