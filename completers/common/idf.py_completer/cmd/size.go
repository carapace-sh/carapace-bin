package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/idf.py_completer/cmd/common"
	"github.com/spf13/cobra"
)

var sizeCmd = &cobra.Command{
	Use:   "size",
	Short: "Print basic size information about the app",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sizeCmd).Standalone()
	common.AddSizeFlags(sizeCmd)
	rootCmd.AddCommand(sizeCmd)
}
