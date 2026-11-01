package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var proplistCmd = &cobra.Command{
	Use:   "proplist",
	Short: "List all properties of a file or directory",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(proplistCmd).Standalone()

	proplistCmd.Flags().IntP("revision", "r", 0, "Refer to a specific revision")
	rootCmd.AddCommand(proplistCmd)

	carapace.Gen(proplistCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
