package goreleaser

import "github.com/carapace-sh/carapace"

// ActionReleaseSteps completes release steps.
//
//	announce
//	archive
func ActionReleaseSteps() carapace.Action {
	return carapace.ActionValues(
		"announce",
		"archive",
		"aur",
		"aur-source",
		"before",
		"chocolatey",
		"docker",
		"homebrew",
		"ko",
		"nfpm",
		"nix",
		"notarize",
		"publish",
		"sbom",
		"scoop",
		"sign",
		"snapcraft",
		"validate",
		"winget",
	).Tag("release steps")
}
