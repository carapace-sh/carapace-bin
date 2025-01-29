package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var context_useCmd = &cobra.Command{
	Use:   "use NAME",
	Short: "Set named destination as default",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(context_useCmd).Standalone()

	contextCmd.AddCommand(context_useCmd)
}
