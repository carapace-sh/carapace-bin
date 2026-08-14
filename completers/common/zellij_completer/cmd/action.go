package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var actionCmd = &cobra.Command{
	Use:     "action",
	Short:   "Send actions to a specific session",
	Aliases: []string{"ac"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(actionCmd).Standalone()

	actionCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(actionCmd)
}
