package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/idf.py_completer/cmd/common"
	"github.com/spf13/cobra"
)

var sizeComponentsCmd = &cobra.Command{
	Use:   "size-components",
	Short: "Print per-component size information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sizeComponentsCmd).Standalone()

	common.AddSizeFlags(sizeComponentsCmd)
	rootCmd.AddCommand(sizeComponentsCmd)
}
