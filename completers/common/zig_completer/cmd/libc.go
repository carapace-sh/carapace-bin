package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var libcCmd = &cobra.Command{
	Use:   "libc",
	Short: "Display native libc paths file or validate one",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(libcCmd).Standalone()

	libcCmd.Flags().StringP("target", "", "", "<arch><sub>-<os>-<abi>")
	libcCmd.Flags().BoolP("help", "h", false, "Print help")

	rootCmd.AddCommand(libcCmd)

	carapace.Gen(libcCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
