package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pnpm"
	"github.com/spf13/cobra"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Create a tarball from a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(packCmd).Standalone()

	packCmd.Flags().Bool("dry-run", false, "Do everything a pack would do except actually creating the tarball")
	packCmd.Flags().String("filter", "", "Restrict the scope to package names matching the given pattern")
	packCmd.Flags().Bool("json", false, "Log output in JSON format")
	packCmd.Flags().String("out", "", "Customize the output tarball path (supports %s for name and %v for version)")
	packCmd.Flags().String("pack-destination", "", "Directory in which pnpm pack will save tarballs")
	packCmd.Flags().String("pack-gzip-level", "", "Custom gzip compression level for the tarball")
	packCmd.Flags().BoolP("recursive", "r", false, "Pack all packages from the workspace")
	rootCmd.AddCommand(packCmd)

	carapace.Gen(packCmd).FlagCompletion(carapace.ActionMap{
		"filter":           pnpm.ActionFilters(),
		"pack-destination": carapace.ActionDirectories(),
	})
}
