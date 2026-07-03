package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var scipCmd = &cobra.Command{
	Use:   "scip",
	Short: "",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(scipCmd).Standalone()

	scipCmd.Flags().String("config-path", "", "A path to an json configuration file that can be used to customize cargo behavior.")
	scipCmd.Flags().Bool("exclude-vendored-libraries", false, "Exclude code from vendored libraries from the resulting index.")
	scipCmd.Flags().String("num-threads", "", "The number of worker threads for cache priming. Defaults to the number of physical cores.")
	scipCmd.Flags().String("output", "", "The output path where the SCIP file will be written to. Defaults to `index.scip`.")
	rootCmd.AddCommand(scipCmd)

	carapace.Gen(scipCmd).PositionalCompletion(
		carapace.ActionFiles(),
	)
}
