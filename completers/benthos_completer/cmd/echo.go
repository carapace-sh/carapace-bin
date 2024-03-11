package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var echoCmd = &cobra.Command{
	Use:   "echo",
	Short: "Parse a config file and echo back a normalised version",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(echoCmd).Standalone()

	rootCmd.AddCommand(echoCmd)
}
