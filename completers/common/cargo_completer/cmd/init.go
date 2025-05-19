package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Create a new cargo package in an existing directory",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groupFor("init"),
}

func init() {
	carapace.Gen(initCmd).Standalone()

	initCmd.Flags().Bool("bin", false, "Use a binary (application) template [default]")
	initCmd.Flags().String("edition", "", "Edition to set for the crate generated")
	initCmd.Flags().BoolP("help", "h", false, "Print help")
	initCmd.Flags().Bool("lib", false, "Use a library template")
	initCmd.Flags().String("name", "", "Set the resulting package name, defaults to the directory name")
	initCmd.Flags().String("registry", "", "Registry to use")
	initCmd.Flags().String("vcs", "", "Initialize a new repository for the given version control system, overriding a global configuration.")
	rootCmd.AddCommand(initCmd)

	carapace.Gen(initCmd).FlagCompletion(carapace.ActionMap{
		"edition":  carapace.ActionValues("2015", "2018"),
		"registry": action.ActionRegistries(),
		"vcs":      carapace.ActionValues("git", "hg", "pijul", "vcs", "none"),
	})

	carapace.Gen(initCmd).PositionalCompletion(
		carapace.ActionDirectories(),
	)
}
