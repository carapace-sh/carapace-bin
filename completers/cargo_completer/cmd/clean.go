package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/completers/cargo_completer/cmd/action"
	"github.com/spf13/cobra"
)

var cleanCmd = &cobra.Command{
	Use:     "clean",
	Short:   "Remove artifacts that cargo has generated in the past",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: groupFor("clean"),
}

func init() {
	carapace.Gen(cleanCmd).Standalone()

	cleanCmd.Flags().Bool("doc", false, "Whether or not to clean just the documentation directory")
	cleanCmd.Flags().BoolP("help", "h", false, "Print help")
	cleanCmd.Flags().String("manifest-path", "", "Path to Cargo.toml")
	cleanCmd.Flags().StringSliceP("package", "p", []string{}, "Package to clean artifacts for")
	cleanCmd.Flags().String("profile", "", "Clean artifacts of the specified profile")
	cleanCmd.Flags().BoolP("quiet", "q", false, "Do not print cargo log messages")
	cleanCmd.Flags().BoolP("release", "r", false, "Whether or not to clean release artifacts")
	cleanCmd.Flags().StringSlice("target", []string{}, "Target triple to clean output for")
	cleanCmd.Flags().String("target-dir", "", "Directory for all generated artifacts")
	rootCmd.AddCommand(cleanCmd)

	carapace.Gen(cleanCmd).FlagCompletion(carapace.ActionMap{
		"manifest-path": carapace.ActionFiles(),
		"package":       action.ActionDependencies(cleanCmd, true),
		"profile":       action.ActionProfiles(cleanCmd),
		"target-dir":    carapace.ActionDirectories(),
	})
}
