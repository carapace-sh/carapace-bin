package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var packCmd = &cobra.Command{
	Use:     "pack",
	Short:   "Generate a tarball from the active workspace",
	GroupID: "general",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(packCmd).Standalone()

	packCmd.Flags().BoolP("dry-run", "n", false, "Print the file paths without actually generating the package archive")
	packCmd.Flags().Bool("install-if-needed", false, "Run a preliminary yarn install if the package contains build scripts")
	packCmd.Flags().Bool("json", false, "Format the output as an NDJSON stream")
	packCmd.Flags().StringP("out", "o", "", "Create the archive at the specified path")
	rootCmd.AddCommand(packCmd)

	carapace.Gen(packCmd).FlagCompletion(carapace.ActionMap{
		"out": carapace.ActionFiles(),
	})
}
