package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var helpmsgCmd = &cobra.Command{
	Use:   "helpmsg",
	Short: "display information about a network message",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(helpmsgCmd).Standalone()
	rootCmd.AddCommand(helpmsgCmd)
}
