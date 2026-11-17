package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var forwardCmd = &cobra.Command{
	Use:   "forward",
	Short: "Forward TCP ports",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(forwardCmd).Standalone()

	rootCmd.AddCommand(forwardCmd)
}
