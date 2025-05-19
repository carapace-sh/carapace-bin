package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/conda"
	"github.com/spf13/cobra"
)

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "Export a given environment",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(exportCmd).Standalone()

	exportCmd.Flags().StringP("channel", "c", "", "Additional channel to include in the export")
	exportCmd.Flags().StringP("file", "f", "", "Environment definition file")
	exportCmd.Flags().Bool("from-history", false, "Build environment spec from explicit specs in history")
	exportCmd.Flags().BoolP("help", "h", false, "Show this help message and exit")
	exportCmd.Flags().Bool("ignore-channels", false, "Do not include channel names with package names.")
	exportCmd.Flags().Bool("json", false, "Report all output as json")
	exportCmd.Flags().StringP("name", "n", "", "Name of environment")
	exportCmd.Flags().Bool("no-builds", false, "Remove build specification from dependencies")
	exportCmd.Flags().Bool("override-channels", false, "Do not include .condarc channels")
	exportCmd.Flags().StringP("prefix", "p", "", "Full path to environment location")
	exportCmd.Flags().BoolP("quiet", "q", false, "Do not display progress bar")
	exportCmd.Flags().BoolP("verbose", "v", false, "Use once for info, twice for debug, three times for trace")
	rootCmd.AddCommand(exportCmd)

	carapace.Gen(exportCmd).FlagCompletion(carapace.ActionMap{
		"file": carapace.ActionFiles(),
		"name": conda.ActionEnvironments(),
	})

}
