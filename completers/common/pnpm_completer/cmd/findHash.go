package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var findHashCmd = &cobra.Command{
	Use:   "find-hash",
	Short: "Experimental! Lists the packages that include the file with the specified hash",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(findHashCmd).Standalone()

	findHashCmd.Flags().BoolP("help", "h", false, "Output usage information")

	rootCmd.AddCommand(findHashCmd)
}
