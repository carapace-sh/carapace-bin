package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add <PATHS>...",
	Short: "Add a new directory or increment its rank",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(addCmd).Standalone()

	addCmd.Flags().BoolP("help", "h", false, "Print help")
	addCmd.Flags().BoolP("version", "V", false, "Print version")
	rootCmd.AddCommand(addCmd)

	carapace.Gen(addCmd).PositionalAnyCompletion(
		carapace.ActionDirectories(),
	)
}
