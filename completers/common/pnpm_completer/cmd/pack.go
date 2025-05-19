package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var packCmd = &cobra.Command{
	Use:   "pack",
	Short: "Create a tarball from a package",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(packCmd).Standalone()

	packCmd.Flags().String("pack-destination", "", "Directory in which `pnpm pack` will save tarballs")
	rootCmd.AddCommand(packCmd)

	carapace.Gen(packCmd).FlagCompletion(carapace.ActionMap{
		"pack-destination": carapace.ActionDirectories(),
	})
}
