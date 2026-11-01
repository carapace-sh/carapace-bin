package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugsslCmd = &cobra.Command{
	Use:    "debugssl",
	Short:  "test a secure connection to a server",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugsslCmd).Standalone()

	rootCmd.AddCommand(debugsslCmd)
}
