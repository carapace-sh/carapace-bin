package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var userInfoCmd = &cobra.Command{
	Use:   "userinfo",
	Short: "Print user information",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(userInfoCmd).Standalone()
	rootCmd.AddCommand(userInfoCmd)
}
