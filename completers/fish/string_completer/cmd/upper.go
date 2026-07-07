package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var upperCmd = &cobra.Command{
	Use:   "upper",
	Short: "Convert strings to uppercase",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(upperCmd).Standalone()

	upperCmd.Flags().BoolP("quiet", "q", false, "suppress output")

	rootCmd.AddCommand(upperCmd)
}
