package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var archimportCmd = &cobra.Command{
	Use:     "archimport",
	Short:   "Import a GNU Arch repository into Git",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groups[group_interaction].ID,
}

func init() {
	carapace.Gen(archimportCmd).Standalone()

	archimportCmd.Flags().StringS("D", "D", "", "follow merge ancestry and attempt to import trees that have been merged from.")
	archimportCmd.Flags().BoolS("T", "T", false, "many tags. Will create a tag for every commit, reflecting the commit name in the Arch repository.")
	archimportCmd.Flags().BoolS("a", "a", false, "attempt to auto-register archives at http://mirrors.sourcecontrol.net")
	archimportCmd.Flags().BoolS("f", "f", false, "use the fast patchset import strategy.")
	archimportCmd.Flags().BoolS("h", "h", false, "display usage.")
	archimportCmd.Flags().BoolS("o", "o", false, "use this for compatibility with old-style branch names used by earlier versions of git archimport.")
	archimportCmd.Flags().StringS("t", "t", "", "override the default tempdir.")
	archimportCmd.Flags().BoolS("v", "v", false, "verbose output.")
	rootCmd.AddCommand(archimportCmd)

	carapace.Gen(archimportCmd).FlagCompletion(carapace.ActionMap{
		"t": carapace.ActionDirectories(),
	})

	// TODO positional completion
	//
	carapace.Gen(archimportCmd).DashAnyCompletion(
		carapace.ActionPositional(archimportCmd),
	)
}
