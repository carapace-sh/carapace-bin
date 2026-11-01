package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var setcomputerpasswordCmd = &cobra.Command{
	Use:   "setcomputerpassword",
	Short: "set the computer account password",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setcomputerpasswordCmd).Standalone()
	rootCmd.AddCommand(setcomputerpasswordCmd)
}
