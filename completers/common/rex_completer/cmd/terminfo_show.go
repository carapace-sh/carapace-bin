package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var terminfoShowCmd = &cobra.Command{
	Use:   "show",
	Short: "Print the terminfo source",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(terminfoShowCmd).Standalone()

	terminfoCmd.AddCommand(terminfoShowCmd)
}
