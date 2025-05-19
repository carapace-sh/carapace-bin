package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/goreleaser"
	"github.com/spf13/cobra"
)

var releaseCmd = &cobra.Command{
	Use:     "release",
	Short:   "Releases the current project",
	Aliases: []string{"r"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(releaseCmd).Standalone()

	releaseCmd.Flags().Bool("auto-snapshot", false, "Automatically sets --snapshot if the repository is dirty")
	releaseCmd.Flags().Bool("clean", false, "Removes the 'dist' directory")
	releaseCmd.Flags().StringP("config", "f", "", "Load configuration from file")
	releaseCmd.Flags().Bool("deprecated", false, "Force print the deprecation message - tests only")
	releaseCmd.Flags().Bool("draft", false, "Whether to set the release to draft. Overrides release.draft in the configuration file")
	releaseCmd.Flags().Bool("fail-fast", false, "Whether to abort the release publishing on the first error")
	releaseCmd.Flags().StringP("parallelism", "p", "", "Amount tasks to run concurrently (default: number of CPUs)")
	releaseCmd.Flags().String("release-footer", "", "Load custom release notes footer from a markdown file")
	releaseCmd.Flags().String("release-footer-tmpl", "", "Load custom release notes footer from a templated markdown file (overrides --release-footer)")
	releaseCmd.Flags().String("release-header", "", "Load custom release notes header from a markdown file")
	releaseCmd.Flags().String("release-header-tmpl", "", "Load custom release notes header from a templated markdown file (overrides --release-header)")
	releaseCmd.Flags().String("release-notes", "", "Load custom release notes from a markdown file (will skip GoReleaser changelog generation)")
	releaseCmd.Flags().String("release-notes-tmpl", "", "Load custom release notes from a templated markdown file (overrides --release-notes)")
	releaseCmd.Flags().Bool("rm-dist", false, "Removes the 'dist' directory")
	releaseCmd.Flags().StringSlice("skip", nil, "Skip the given options (valid options are announce, aur, before, chocolatey, docker, homebrew, ko, nfpm, nix, notarize, publish, sbom, scoop, sign, snapcraft, validate, winget)")
	releaseCmd.Flags().Bool("skip-announce", false, "Skips announcing releases (implies --skip=validate)")
	releaseCmd.Flags().Bool("skip-before", false, "Skips global before hooks")
	releaseCmd.Flags().Bool("skip-docker", false, "Skips Docker Images/Manifests builds")
	releaseCmd.Flags().Bool("skip-ko", false, "Skips Ko builds")
	releaseCmd.Flags().Bool("skip-publish", false, "Skips publishing artifacts (implies --skip=announce)")
	releaseCmd.Flags().Bool("skip-sbom", false, "Skips cataloging artifacts")
	releaseCmd.Flags().Bool("skip-sign", false, "Skips signing artifacts")
	releaseCmd.Flags().Bool("skip-validate", false, "Skips git checks")
	releaseCmd.Flags().Bool("snapshot", false, "Generate an unversioned snapshot release, skipping all validations and without publishing any artifacts (implies --skip=announce,publish,validate)")
	releaseCmd.Flags().String("timeout", "", "Timeout to the entire release process")
	releaseCmd.Flag("deprecated").Hidden = true
	releaseCmd.Flag("rm-dist").Hidden = true
	releaseCmd.Flag("skip-announce").Hidden = true
	releaseCmd.Flag("skip-before").Hidden = true
	releaseCmd.Flag("skip-docker").Hidden = true
	releaseCmd.Flag("skip-ko").Hidden = true
	releaseCmd.Flag("skip-publish").Hidden = true
	releaseCmd.Flag("skip-sbom").Hidden = true
	releaseCmd.Flag("skip-sign").Hidden = true
	releaseCmd.Flag("skip-validate").Hidden = true
	rootCmd.AddCommand(releaseCmd)

	carapace.Gen(releaseCmd).FlagCompletion(carapace.ActionMap{
		"config":              carapace.ActionFiles(),
		"release-footer":      carapace.ActionFiles(),
		"release-footer-tmpl": carapace.ActionFiles(),
		"release-header":      carapace.ActionFiles(),
		"release-header-tmpl": carapace.ActionFiles(),
		"release-notes":       carapace.ActionFiles(),
		"release-notes-tmpl":  carapace.ActionFiles(),
		"skip":                goreleaser.ActionReleaseSteps().UniqueList(","),
	})
}
