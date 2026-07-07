package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var split0Cmd = &cobra.Command{
	Use:   "split0",
	Short: "Split strings on NUL bytes",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(split0Cmd).Standalone()

	split0Cmd.Flags().BoolP("allow-empty", "a", false, "print nothing for missing fields")
	split0Cmd.Flags().StringP("fields", "f", "", "print specific fields")
	split0Cmd.Flags().StringP("max", "m", "", "maximum number of splits")
	split0Cmd.Flags().BoolP("no-empty", "n", false, "exclude empty results")
	split0Cmd.Flags().BoolP("quiet", "q", false, "suppress output")
	split0Cmd.Flags().BoolP("right", "r", false, "split right-to-left")

	rootCmd.AddCommand(split0Cmd)
}
