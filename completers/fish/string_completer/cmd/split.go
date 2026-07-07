package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var splitCmd = &cobra.Command{
	Use:   "split",
	Short: "Split strings on a separator",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(splitCmd).Standalone()

	splitCmd.Flags().BoolP("allow-empty", "a", false, "print nothing for missing fields")
	splitCmd.Flags().StringP("fields", "f", "", "print specific fields")
	splitCmd.Flags().StringP("max", "m", "", "maximum number of splits")
	splitCmd.Flags().BoolP("no-empty", "n", false, "exclude empty results")
	splitCmd.Flags().BoolP("quiet", "q", false, "suppress output")
	splitCmd.Flags().BoolP("right", "r", false, "split right-to-left")

	rootCmd.AddCommand(splitCmd)
}
