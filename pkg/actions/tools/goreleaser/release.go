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
		"changelog",
		"chocolatey",
		"docker",
		"flatpak",
		"homebrew",
		"iru",
		"ko",
		"makeself",
		"mcp",
		"nfpm",
		"nix",
		"notarize",
		"publish",
		"sbom",
		"scoop",
		"sign",
		"snapcraft",
		"srpm",
		"validate",
		"winget",
	).Tag("release steps")
}
