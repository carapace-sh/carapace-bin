package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var objcopyCmd = &cobra.Command{
	Use:   "objcopy",
	Short: "Copy and translate object files",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(objcopyCmd).Standalone()

	objcopyCmd.Flags().StringP("output-target", "O", "", "Format of the output file")
	objcopyCmd.Flags().StringP("only-section", "j", "", "Remove all but <section>")
	objcopyCmd.Flags().Bool("strip-debug", false, "Remove debugging symbols only")
	objcopyCmd.Flags().Bool("strip-all", false, "Remove all symbols")
	objcopyCmd.Flags().Bool("only-keep-debug", false, "Strip a file, but keep debug sections")
	objcopyCmd.Flags().BoolP("help", "h", false, "Print help")

	rootCmd.AddCommand(objcopyCmd)

	carapace.Gen(objcopyCmd).PositionalAnyCompletion(
		carapace.ActionFiles(),
	)
}
