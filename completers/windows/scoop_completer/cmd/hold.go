package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var holdCmd = &cobra.Command{
	Use:   "hold",
	Short: "hold an app to disable updates",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(holdCmd).Standalone()
	holdCmd.Flags().BoolP("global", "g", false, "hold globally installed apps")
	rootCmd.AddCommand(holdCmd)
}
