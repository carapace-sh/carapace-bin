package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rubCmd = &cobra.Command{
	Use:   "rub",
	Short: "Combines two entities together to perform an operation",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rubCmd).Standalone()

	rubCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	rootCmd.AddCommand(rubCmd)
}
