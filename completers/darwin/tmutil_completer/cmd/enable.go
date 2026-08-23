package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(enableCmd)
	rootCmd.AddCommand(disableCmd)
}

var enableCmd = &cobra.Command{
	Use:   "enable",
	Short: "turn on automatic backups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "turn off automatic backups",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(enableCmd).Standalone()
	carapace.Gen(disableCmd).Standalone()
}