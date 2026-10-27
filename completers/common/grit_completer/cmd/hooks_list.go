package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var hooks_listCmd = &cobra.Command{
	Use:   "list",
	Short: "List hook scripts in the resolved hooks directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(hooks_listCmd).Standalone()

	hooks_listCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	hooks_listCmd.Flags().StringP("limit", "n", "", "Maximum number of rows to list")
	hooks_listCmd.Flags().String("max-count", "", "Maximum number of rows to list")
	hooksCmd.AddCommand(hooks_listCmd)
}
