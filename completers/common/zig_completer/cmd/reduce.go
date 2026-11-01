package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var reduceCmd = &cobra.Command{
	Use:     "reduce [options] ./checker root_source_file.zig [-- [argv]]",
	Short:   "Minimize a bug report",
	GroupID: "source",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reduceCmd).Standalone()

	rootCmd.AddCommand(reduceCmd)

	reduceCmd.Flags().StringSlice("deps", nil, "Set dependency names for the root package")
	reduceCmd.Flags().BoolP("help", "h", false, "Print help")
	reduceCmd.Flags().String("main-mod-path", "", "Set the directory of the root module")
	reduceCmd.Flags().StringSlice("mod", nil, "Make a module available for dependency under the given name")
	reduceCmd.Flags().String("seed", "", "Override the random seed (default: 0)")
	reduceCmd.Flags().Bool("skip-smoke-test", false, "Skip interestingness check smoke test")

	carapace.Gen(reduceCmd).PositionalCompletion(
		carapace.ActionFiles(),
		carapace.ActionFiles(".zig"),
	)
}
