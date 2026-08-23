package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var waitCmd = &cobra.Command{
	Use:   "wait",
	Short: "Wait before executing the next command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(waitCmd).Standalone()
	waitCmd.Flags().String("duration", "", "Duration in seconds")
	waitCmd.Flags().String("file", "", "Wait until file exists")
	waitCmd.Flags().String("gone", "", "Wait until file no longer exists")
	rootCmd.AddCommand(waitCmd)
}
