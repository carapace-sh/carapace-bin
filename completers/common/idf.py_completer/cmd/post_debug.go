package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var postDebugCmd = &cobra.Command{
	Use:   "post-debug",
	Short: "Run post-debug actions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(postDebugCmd).Standalone()
	postDebugCmd.Flags().String("block", "", "Set to 1 for blocking console on outputs of async debug actions")
	rootCmd.AddCommand(postDebugCmd)
}
