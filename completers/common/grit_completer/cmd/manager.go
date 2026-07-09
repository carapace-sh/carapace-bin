package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var managerCmd = &cobra.Command{
	Use:   "manager",
	Short: "Git credential helper backed by the Windows Credential Manager",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(managerCmd).Standalone()

	managerCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(managerCmd)

	carapace.Gen(managerCmd).PositionalCompletion(
		carapace.ActionValues("get", "store", "erase"),
	)
}
