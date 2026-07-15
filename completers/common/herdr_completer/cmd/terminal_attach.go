package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var terminal_attachCmd = &cobra.Command{
	Use:   "attach <terminal_id>",
	Short: "Attach directly to a terminal stream",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(terminal_attachCmd).Standalone()

	terminal_attachCmd.Flags().Bool("takeover", false, "")
	terminalCmd.AddCommand(terminal_attachCmd)
}
