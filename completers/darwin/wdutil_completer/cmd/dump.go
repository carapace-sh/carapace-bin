package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dumpCmd = &cobra.Command{
	Use:   "dump",
	Short: "dump Wi-Fi log buffer to file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dumpCmd).Standalone()
	rootCmd.AddCommand(dumpCmd)
}