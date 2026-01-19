package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debug_getCmd = &cobra.Command{
	Use:   "get",
	Short: "retrieve a blob from contentstore",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debug_getCmd).Standalone()

	debugCmd.AddCommand(debug_getCmd)
}
