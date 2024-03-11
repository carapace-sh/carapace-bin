package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/pacman"
	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Clone or sync packages build files from AUR",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloneCmd).Standalone()

	cloneCmd.Flags().String("builddir", "", "build directory")
	cloneCmd.Flags().Bool("overwrite", false, "overwrite existing files")
	cloneCmd.Flags().BoolP("quiet", "q", false, "do not print any output")
	cloneCmd.Flags().BoolP("recurse", "r", false, "also clone needed dependencies")
	rootCmd.AddCommand(cloneCmd)

	carapace.Gen(cloneCmd).FlagCompletion(carapace.ActionMap{
		"builddir": carapace.ActionDirectories(),
	})

	carapace.Gen(cloneCmd).PositionalAnyCompletion(
		pacman.ActionPackages().FilterArgs(),
	)

}
