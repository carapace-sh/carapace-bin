package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var executeWithPrivilegesCmd = &cobra.Command{
	Use:   "execute-with-privileges",
	Short: "Execute tool with privileges",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(executeWithPrivilegesCmd).Standalone()
	rootCmd.AddCommand(executeWithPrivilegesCmd)
}
