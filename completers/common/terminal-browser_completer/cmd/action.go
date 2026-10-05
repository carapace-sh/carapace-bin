package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var actionCmd = &cobra.Command{
	Use:   "action",
	Short: "Use the open browser through the agent-browser CLI",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(actionCmd).Standalone()

	actionCmd.Flags().String("browser", "", "A browser key from terminal-browser ls")
	actionCmd.Flags().Bool("follow", false, "Bring the tab to the front before running the command")
	actionCmd.Flags().String("tab", "", "A tab id from terminal-browser ls")
	actionCmd.Flags().String("target", "", "A CDP target id")
	rootCmd.AddCommand(actionCmd)
}
