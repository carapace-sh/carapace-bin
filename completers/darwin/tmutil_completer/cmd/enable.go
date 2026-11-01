package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "turn on automatic backups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(enableCmd).Standalone()
	rootCmd.AddCommand(enableCmd)
}
