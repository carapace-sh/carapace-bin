package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var promisorCmd = &cobra.Command{
	Use:    "promisor",
	Short:  "Walk remote history and index keys as promisor entries",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(promisorCmd).Standalone()

	promisorCmd.Flags().BoolP("help", "h", false, "Print help")
	rootCmd.AddCommand(promisorCmd)
}
