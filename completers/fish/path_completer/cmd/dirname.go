package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dirnameCmd = &cobra.Command{
	Use:   "dirname",
	Short: "Return the directory portion of the path",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dirnameCmd).Standalone()

	dirnameCmd.Flags().BoolS("Z", "Z", false, "NUL-delimited output")
	dirnameCmd.Flags().Bool("null-in", false, "NUL-delimited input")
	dirnameCmd.Flags().Bool("null-out", false, "NUL-delimited output")
	dirnameCmd.Flags().BoolS("q", "q", false, "suppress output")
	dirnameCmd.Flags().Bool("quiet", false, "suppress output")
	dirnameCmd.Flags().BoolS("z", "z", false, "NUL-delimited input")

	rootCmd.AddCommand(dirnameCmd)
}
