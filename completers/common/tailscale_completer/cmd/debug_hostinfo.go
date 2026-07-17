package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_hostinfoCmd = &cobra.Command{
	Use:   "hostinfo",
	Short: "Print hostinfo",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_hostinfoCmd).Standalone()

	debugCmd.AddCommand(debug_hostinfoCmd)
}
