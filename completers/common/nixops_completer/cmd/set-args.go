package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var SetArgsCmd = &cobra.Command{
	Use:   "set-args",
	Short: "Set Args",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(SetArgsCmd).Standalone()
	rootCmd.AddCommand(SetArgsCmd)
}
