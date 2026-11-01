package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var behaviorCmd = &cobra.Command{
	Use:   "behavior",
	Short: "query and set volume behavior",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(behaviorCmd).Standalone()
	rootCmd.AddCommand(behaviorCmd)
}
