package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sessionCmd = &cobra.Command{
	Use:   "session",
	Short: "Generate random session key",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sessionCmd).Standalone()

	sessionCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(sessionCmd)
}
