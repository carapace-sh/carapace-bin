package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var help_findHashCmd = &cobra.Command{
	Use:   "find-hash",
	Short: "Lists the packages that include the file with the specified hash",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(help_findHashCmd).Standalone()

	helpCmd.AddCommand(help_findHashCmd)
}
