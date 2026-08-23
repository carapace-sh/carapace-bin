package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "enable or disable logging",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logCmd).Standalone()
	logCmd.Flags().Bool("system", false, "Enable/disable system logging")
	logCmd.Flags().Bool("wifi", false, "Enable/disable Wi-Fi logging")
	rootCmd.AddCommand(logCmd)
}
