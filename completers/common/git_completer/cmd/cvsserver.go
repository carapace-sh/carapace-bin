package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cvsserverCmd = &cobra.Command{
	Use:     "cvsserver",
	Short:   "A CVS server emulator for Git",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_interaction].ID,
}

func init() {
	carapace.Gen(cvsserverCmd).Standalone()

	cvsserverCmd.Flags().String("base-path", "", "prepend path to requested CVSROOT")
	cvsserverCmd.Flags().Bool("export-all", false, "don’t check for gitcvs.enabled in config ")
	cvsserverCmd.Flags().BoolP("help", "h", false, "print usage information and exit")
	cvsserverCmd.Flags().Bool("strict-paths", false, "don’t allow recursing into subdirectories")
	cvsserverCmd.Flags().BoolP("version", "V", false, "print version information and exit")
	rootCmd.AddCommand(cvsserverCmd)

	carapace.Gen(cvsserverCmd).PositionalAnyCompletion(
		carapace.ActionDirectories(),
	)
}
