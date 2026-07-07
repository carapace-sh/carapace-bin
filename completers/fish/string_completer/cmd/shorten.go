package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shortenCmd = &cobra.Command{
	Use:   "shorten",
	Short: "Truncate strings to a visible width",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shortenCmd).Standalone()

	shortenCmd.Flags().BoolS("N", "N", false, "only use first line of each string")
	shortenCmd.Flags().StringP("char", "c", "", "character(s) to use as ellipsis")
	shortenCmd.Flags().BoolP("left", "l", false, "remove text from the left")
	shortenCmd.Flags().StringP("max", "m", "", "maximum visible width")
	shortenCmd.Flags().Bool("no-newline", false, "only use first line of each string")
	shortenCmd.Flags().BoolP("quiet", "q", false, "suppress output")

	rootCmd.AddCommand(shortenCmd)
}
