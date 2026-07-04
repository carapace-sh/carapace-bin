package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var livecheckCmd = &cobra.Command{
	Use:   "livecheck",
	Short: "Check if software has been updated since the Portfile was last modified",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(livecheckCmd).Standalone()
	rootCmd.AddCommand(livecheckCmd)
}
