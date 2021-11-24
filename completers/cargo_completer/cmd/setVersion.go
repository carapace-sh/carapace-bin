package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var setVersionCmd = &cobra.Command{
	Use:   "set-version",
	Short: "Change a package's version in the local manifest file",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(setVersionCmd).Standalone()

	setVersionCmd.Flags().Bool("all", false, "[deprecated in favor of `--workspace`]")
	setVersionCmd.Flags().String("bump", "", "Increment manifest version")
	setVersionCmd.Flags().Bool("dry-run", false, "Print changes to be made without making them")
	setVersionCmd.Flags().String("exclude", "", "Crates to exclude and not modify")
	setVersionCmd.Flags().BoolP("help", "h", false, "Prints help information")
	setVersionCmd.Flags().String("manifest-path", "", "Path to the manifest to upgrade")
	setVersionCmd.Flags().StringP("metadata", "m", "", "Specify the version metadata field (e.g. a wrapped libraries version)")
	setVersionCmd.Flags().StringP("package", "p", "", "Package id of the crate to change the version of")
	setVersionCmd.Flags().BoolP("version", "V", false, "Prints version information")
	setVersionCmd.Flags().Bool("workspace", false, "Modify all packages in the workspace")
	rootCmd.AddCommand(setVersionCmd)

	// TODO flag completion
	carapace.Gen(setVersionCmd).FlagCompletion(carapace.ActionMap{
		"bump":         carapace.ActionValues("major", "minor", "patch", "release", "rc", "beta", "alpha"),
		"manifes-path": carapace.ActionFiles(".toml"),
	})

	// TODO complete current version as positional for easy edit
}
