package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var killSessionCmd = &cobra.Command{
	Use:   "kill-session",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(killSessionCmd).Standalone()

	killSessionCmd.Flags().StringS("t", "t", "", "target-session")
	rootCmd.AddCommand(killSessionCmd)
}
