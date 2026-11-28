package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/idf.py_completer/cmd/common"
	"github.com/spf13/cobra"
)

var sizeFilesCmd = &cobra.Command{
	Use:   "size-files",
	Short: "Print per-source-file size information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sizeFilesCmd).Standalone()
	common.AddSizeFlags(sizeFilesCmd)
	rootCmd.AddCommand(sizeFilesCmd)
}
