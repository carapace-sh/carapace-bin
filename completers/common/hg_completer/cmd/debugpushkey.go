package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugpushkeyCmd = &cobra.Command{
	Use:    "debugpushkey",
	Short:  "REPO NAMESPACE [KEY OLD NEW]",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugpushkeyCmd).Standalone()

	rootCmd.AddCommand(debugpushkeyCmd)
}
