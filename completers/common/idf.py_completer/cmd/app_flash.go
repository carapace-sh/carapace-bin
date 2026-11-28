package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/idf.py_completer/cmd/common"
	"github.com/spf13/cobra"
)

var appFlashCmd = &cobra.Command{
	Use:   "app-flash",
	Short: "Flash the app only",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appFlashCmd).Standalone()

	common.AddFlashFlags(appFlashCmd)
	rootCmd.AddCommand(appFlashCmd)
}
